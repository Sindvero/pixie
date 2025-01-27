/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include "src/carnot/planner/logical_planner.h"

#include <utility>

#include "src/carnot/planner/compiler_state/compiler_state.h"
#include "src/carnot/planner/ir/ast_utils.h"
#include "src/carnot/planner/parser/parser.h"
#include "src/shared/scriptspb/scripts.pb.h"

namespace px {
namespace carnot {
namespace planner {

using table_store::schemapb::Schema;

StatusOr<std::unique_ptr<RelationMap>> MakeRelationMapFromSchema(const Schema& schema_pb) {
  auto rel_map = std::make_unique<RelationMap>();
  for (auto& relation_pair : schema_pb.relation_map()) {
    px::table_store::schema::Relation rel;
    PL_RETURN_IF_ERROR(rel.FromProto(&relation_pair.second));
    rel_map->emplace(relation_pair.first, rel);
  }

  return rel_map;
}
StatusOr<std::unique_ptr<RelationMap>> MakeRelationMapFromDistributedState(
    const distributedpb::DistributedState& state_pb) {
  auto rel_map = std::make_unique<RelationMap>();
  for (const auto& schema_info : state_pb.schema_info()) {
    px::table_store::schema::Relation rel;
    PL_RETURN_IF_ERROR(rel.FromProto(&schema_info.relation()));
    rel_map->emplace(schema_info.name(), rel);
  }

  return rel_map;
}

static inline RedactionOptions RedactionOptionsFromPb(
    const distributedpb::RedactionOptions& redaction_options) {
  RedactionOptions options;
  options.use_full_redaction = redaction_options.use_full_redaction();
  options.use_px_redact_pii_best_effort = redaction_options.use_px_redact_pii_best_effort();
  return options;
}

StatusOr<std::unique_ptr<CompilerState>> CreateCompilerState(
    const distributedpb::LogicalPlannerState& logical_state, RegistryInfo* registry_info,
    int64_t max_output_rows_per_table) {
  PL_ASSIGN_OR_RETURN(std::unique_ptr<RelationMap> rel_map,
                      MakeRelationMapFromDistributedState(logical_state.distributed_state()));

  SensitiveColumnMap sensitive_columns = {
      {"cql_events", {"req_body", "resp_body"}},
      {"http_events", {"req_headers", "req_body", "resp_headers", "resp_body"}},
      {"kafka_events.beta", {"req_body", "resp"}},
      {"mysql_events", {"req_body", "resp_body"}},
      {"nats_events.beta", {"body", "resp"}},
      {"pgsql_events", {"req", "resp"}},
      {"redis_events", {"req_args", "resp"}}};

  std::unique_ptr<planpb::OTelEndpointConfig> otel_endpoint_config = nullptr;
  if (logical_state.has_otel_endpoint_config()) {
    otel_endpoint_config = std::make_unique<planpb::OTelEndpointConfig>();
    otel_endpoint_config->set_url(logical_state.otel_endpoint_config().url());
    for (const auto& [key, value] : logical_state.otel_endpoint_config().headers()) {
      (*otel_endpoint_config->mutable_headers())[key] = value;
    }
    otel_endpoint_config->set_insecure(logical_state.otel_endpoint_config().insecure());
  }
  std::unique_ptr<planner::PluginConfig> plugin_config = nullptr;
  if (logical_state.has_plugin_config()) {
    plugin_config = std::unique_ptr<planner::PluginConfig>(
        new planner::PluginConfig{logical_state.plugin_config().start_time_ns(),
                                  logical_state.plugin_config().end_time_ns()});
  }
  planner::DebugInfo debug_info;
  for (const auto& debug_info_pb : logical_state.debug_info().otel_debug_attributes()) {
    debug_info.otel_debug_attrs.push_back({debug_info_pb.name(), debug_info_pb.value()});
  }
  // Create a CompilerState obj using the relation map and grabbing the current time.
  return std::make_unique<planner::CompilerState>(
      std::move(rel_map), sensitive_columns, registry_info, px::CurrentTimeNS(),
      max_output_rows_per_table, logical_state.result_address(),
      logical_state.result_ssl_targetname(),
      // TODO(philkuz) add an endpoint config to logical_state and pass that in here.
      RedactionOptionsFromPb(logical_state.redaction_options()), std::move(otel_endpoint_config),
      // TODO(philkuz) propagate the otel debug attributes here.
      std::move(plugin_config), debug_info);
}

StatusOr<std::unique_ptr<LogicalPlanner>> LogicalPlanner::Create(const udfspb::UDFInfo& udf_info) {
  auto planner = std::unique_ptr<LogicalPlanner>(new LogicalPlanner());
  PL_RETURN_IF_ERROR(planner->Init(udf_info));
  return planner;
}

Status LogicalPlanner::Init(const udfspb::UDFInfo& udf_info) {
  compiler_ = compiler::Compiler();
  registry_info_ = std::make_unique<planner::RegistryInfo>();
  PL_RETURN_IF_ERROR(registry_info_->Init(udf_info));

  PL_ASSIGN_OR_RETURN(distributed_planner_, distributed::DistributedPlanner::Create());
  return Status::OK();
}

StatusOr<std::unique_ptr<distributed::DistributedPlan>> LogicalPlanner::Plan(
    const distributedpb::LogicalPlannerState& logical_state,
    const plannerpb::QueryRequest& query_request) {
  // Compile into the IR.
  auto ms = logical_state.plan_options().max_output_rows_per_table();
  VLOG(1) << "Max output rows: " << ms;
  PL_ASSIGN_OR_RETURN(std::unique_ptr<CompilerState> compiler_state,
                      CreateCompilerState(logical_state, registry_info_.get(), ms));

  std::vector<plannerpb::FuncToExecute> exec_funcs(query_request.exec_funcs().begin(),
                                                   query_request.exec_funcs().end());
  PL_ASSIGN_OR_RETURN(
      std::shared_ptr<IR> single_node_plan,
      compiler_.CompileToIR(query_request.query_str(), compiler_state.get(), exec_funcs));
  // Create the distributed plan.
  PL_ASSIGN_OR_RETURN(auto distributed_plan,
                      distributed_planner_->Plan(logical_state.distributed_state(),
                                                 compiler_state.get(), single_node_plan.get()));
  distributed_plan->SetExecutionCompleteAddress(logical_state.result_address(),
                                                logical_state.result_ssl_targetname());
  return distributed_plan;
}

StatusOr<std::unique_ptr<compiler::MutationsIR>> LogicalPlanner::CompileTrace(
    const distributedpb::LogicalPlannerState& logical_state,
    const plannerpb::CompileMutationsRequest& mutations_req) {
  // Compile into the IR.
  auto ms = logical_state.plan_options().max_output_rows_per_table();
  VLOG(1) << "Max output rows: " << ms;
  PL_ASSIGN_OR_RETURN(std::unique_ptr<CompilerState> compiler_state,
                      CreateCompilerState(logical_state, registry_info_.get(), ms));

  std::vector<plannerpb::FuncToExecute> exec_funcs(mutations_req.exec_funcs().begin(),
                                                   mutations_req.exec_funcs().end());

  return compiler_.CompileTrace(mutations_req.query_str(), compiler_state.get(), exec_funcs);
}

StatusOr<absl::flat_hash_map<std::string, ::px::table_store::schemapb::Relation>>
LogicalPlanner::CalculateOutputSchemas(const distributedpb::LogicalPlannerState& logical_state,
                                       const std::string& pxl_script) {
  PL_ASSIGN_OR_RETURN(
      std::unique_ptr<CompilerState> compiler_state,
      CreateCompilerState(logical_state, registry_info_.get(), /* max_output_rows */ 0));

  PL_ASSIGN_OR_RETURN(std::shared_ptr<IR> single_node_plan,
                      compiler_.CompileToIR(pxl_script, compiler_state.get(), {}));

  absl::flat_hash_map<std::string, ::px::table_store::schemapb::Relation> output_schemas;
  for (const auto& n : single_node_plan->FindNodesThatMatch(ExternalGRPCSink())) {
    auto gsink = static_cast<GRPCSinkIR*>(n);
    PL_ASSIGN_OR_RETURN(auto relation, gsink->resolved_table_type()->ToRelation());
    table_store::schemapb::Relation* relation_pb = &output_schemas[gsink->name()];
    PL_RETURN_IF_ERROR(relation.ToProto(relation_pb));
  }
  return output_schemas;
}

StatusOr<std::string> LogicalPlanner::GetUnusedVarName(
    const distributedpb::LogicalPlannerState& logical_state, const std::string& script,
    const std::string& base_name) const {
  Parser parser;
  PL_ASSIGN_OR_RETURN(pypa::AstModulePtr ast, parser.Parse(script));

  bool func_based_exec = false;
  absl::flat_hash_set<std::string> reserved_names;
  compiler::ModuleHandler module_handler;
  compiler::MutationsIR mutations_ir;
  std::shared_ptr<IR> ir = std::make_shared<IR>();
  auto var_table = compiler::VarTable::Create();
  PL_ASSIGN_OR_RETURN(
      std::unique_ptr<CompilerState> compiler_state,
      CreateCompilerState(logical_state, registry_info_.get(), /* max_output_rows */ 0));
  PL_ASSIGN_OR_RETURN(auto ast_walker,
                      compiler::ASTVisitorImpl::Create(
                          ir.get(), var_table, &mutations_ir, compiler_state.get(), &module_handler,
                          func_based_exec, absl::flat_hash_set<std::string>{}));

  PL_RETURN_IF_ERROR(ast_walker->ProcessModuleNode(ast));
  auto cur_name = base_name;
  int64_t counter = 0;
  while (var_table->HasVariable(cur_name)) {
    if (counter > 1000) {
      return error::InvalidArgument("Gave up searching for an unused variable name with base: $0",
                                    base_name);
    }
    cur_name = absl::Substitute("$0_$1", base_name, counter);
    ++counter;
  }

  return cur_name;
}

StatusOr<std::vector<LogicalPlanner::DisplayLine>> LogicalPlanner::GetPxDisplayLines(
    const std::string& script) {
  // Parse the script into an ast.
  // Check for any calls to px.display().
  // Make sure the arguments are expected and valid.
  Parser parser;
  PL_ASSIGN_OR_RETURN(pypa::AstModulePtr ast, parser.Parse(script));
  std::vector<LogicalPlanner::DisplayLine> display_lines;

  std::vector<std::string> all_script_lines = absl::StrSplit(script, '\n');

  // We are strictly looking for px.display calls here.
  for (const auto& [i, stmt] : Enumerate(ast->body->items)) {
    if (stmt->type != pypa::AstType::ExpressionStatement) {
      continue;
    }
    auto expr = PYPA_PTR_CAST(ExpressionStatement, stmt)->expr;
    if (expr->type != pypa::AstType::Call) {
      continue;
    }
    auto call = PYPA_PTR_CAST(Call, expr);
    // We check if the function is an attribute px.display.
    if (call->function->type != pypa::AstType::Attribute) {
      continue;
    }
    auto function = PYPA_PTR_CAST(Attribute, call->function);
    if (function->value->type != pypa::AstType::Name) {
      continue;
    }

    auto fn_value = PYPA_PTR_CAST(Name, function->value);
    if (fn_value->id != "px") {
      continue;
    }

    if (function->attribute->type != pypa::AstType::Name) {
      continue;
    }

    auto fn_attribute = PYPA_PTR_CAST(Name, function->attribute);
    if (fn_attribute->id != "display") {
      continue;
    }
    // Everything after this will set expectations for the arguments for px.display.
    // If anything is not as expected, we will return an error instead of skipping.

    // We expect two arguments, the first being a dataframe expression, and the second being a
    // string.

    if (call->arguments.size() != 2) {
      return CreateAstError(call, "expected two arguments to px.display, got $0",
                            call->arguments.size());
    }

    if (call->arguments[1]->type != pypa::AstType::Str) {
      return CreateAstError(call->arguments[1],
                            "expected second argument to px.display to be a string, received a $0",
                            GetAstTypeName(call->arguments[1]->type));
    }

    auto first_line = stmt->line - 1;
    int64_t last_line;
    if (i == ast->body->items.size() - 1) {
      last_line = static_cast<int64_t>(all_script_lines.size()) - 1;
    } else {
      // Here we assign the last line to be the line before the next statement.
      // We subtract 2 from the line number of the next statement for the following reasons:
      // Ast line numbers are 1-indexed (first line is 1), but GetPxDisplayLines is 0-indexed.
      // So we first have to subtract 1 to convert 1-index to 0-index, then we
      // subtract 1 again to get the line before the next statement.
      last_line = ast->body->items[i + 1]->line - 2;
    }

    auto table_name = PYPA_PTR_CAST(Str, call->arguments[1])->value;

    // Somehow parse this from the string.
    PL_ASSIGN_OR_RETURN(auto dataframe_argument, AstToString(call->arguments[0]));

    auto statement_str = absl::StrJoin(all_script_lines.begin() + first_line,
                                       all_script_lines.begin() + last_line + 1, "\n");

    display_lines.push_back(LogicalPlanner::DisplayLine{
        statement_str,
        table_name,
        dataframe_argument,
        first_line,
        last_line,
    });
  }

  return display_lines;
}

std::string IndentBlock(const std::string& block, const std::string& indent) {
  return indent + absl::StrJoin(absl::StrSplit(block, '\n'), "\n" + indent);
}

const char kGaugeFormat[] = R"pxl(px.otel.metric.Gauge(
  name='$0',
  description='',
  value=$1,
))pxl";

StatusOr<std::string> RelationToOTelExport(const std::string& table_name,
                                           const std::string& unique_df_name,
                                           const px::table_store::schemapb::Relation& relation) {
  std::vector<std::string> resource_fields;
  std::vector<std::string> data_exports;
  std::string service_col = "";
  bool has_time_column = false;
  for (const auto& column : relation.columns()) {
    if (column.column_semantic_type() == ::px::types::ST_DURATION_NS_QUANTILES ||
        column.column_semantic_type() == ::px::types::ST_QUANTILES) {
      return error::InvalidArgument(
          "quantiles are not supported yet for generation of OTel export scripts");
    }
    if (column.column_name() == "time_") {
      if (column.column_type() != ::px::types::TIME64NS) {
        return error::InvalidArgument("time_ column must be of type TIME64NS, received $0",
                                      ToString(column.column_type()));
      }
      // The export script expects a time column to exist, but does not specify it.
      has_time_column = true;
      continue;
    }

    std::string name = table_name + "." + column.column_name();
    std::string df_col = unique_df_name + "." + column.column_name();

    if (column.column_semantic_type() == ::px::types::ST_SERVICE_NAME) {
      if (service_col == "" || column.column_name() == "service") {
        service_col = df_col;
      }
    }
    switch (column.column_type()) {
      case ::px::types::BOOLEAN:
      case ::px::types::STRING: {
        resource_fields.push_back(absl::Substitute("'$0': $1", name, df_col));
        break;
      }
      case ::px::types::INT64:
      case ::px::types::FLOAT64:
        data_exports.push_back(absl::Substitute(kGaugeFormat, name, df_col));
        break;
      case ::px::types::UINT128:
        return error::InvalidArgument(
            "column '$0' uses an unsupported type: UINT128. Please convert the column to a string",
            column.column_name());
      case ::px::types::TIME64NS:
        return error::InvalidArgument(
            "illegal column '$0' -> TIME64NS column not named 'time_' is ambiguous. Please file a "
            "feature request on GitHub if you have a clear use case for TIME64NS columns",
            column.column_name());
      case ::px::types::DATA_TYPE_UNKNOWN:
      case ::px::types::DataType_INT_MIN_SENTINEL_DO_NOT_USE_:
      case ::px::types::DataType_INT_MAX_SENTINEL_DO_NOT_USE_:
        return error::InvalidArgument("unsupported type $0", ToString(column.column_type()));
    }
  }
  if (!has_time_column) {
    return error::InvalidArgument("Table '$0' does not have a time_ column of TIME64NS type",
                                  table_name);
  }

  if (service_col == "") {
    return error::InvalidArgument(
        "Table '$0' does not have a service column. Make sure you create a service column ie "
        "`df.ctx['service']` and include it in any groupbys and joins",
        table_name);
  }

  if (data_exports.empty()) {
    return error::InvalidArgument(
        "Table '$0' does not have any INT64 or FLOAT64 that can be converted to OTel metrics",
        table_name);
  }

  resource_fields.push_back(absl::Substitute("'service.name': $0", service_col));

  std::string body = absl::Substitute("resource={\n$0\n},\ndata=[\n$1\n]",
                                      IndentBlock(absl::StrJoin(resource_fields, ",\n"), "  "),
                                      IndentBlock(absl::StrJoin(data_exports, ",\n"), "  "));
  return absl::Substitute("px.export($0, px.otel.Data(\n$1\n))", unique_df_name,
                          IndentBlock(body, "  "));
}

StatusOr<std::unique_ptr<plannerpb::GenerateOTelScriptResponse>> LogicalPlanner::GenerateOTelScript(
    const plannerpb::GenerateOTelScriptRequest& generate_req) {
  PL_ASSIGN_OR_RETURN(auto schema, CalculateOutputSchemas(generate_req.logical_planner_state(),
                                                          generate_req.pxl_script()));
  if (schema.size() == 0) {
    return error::InvalidArgument("script does not have any output tables");
  }

  PL_ASSIGN_OR_RETURN(auto display_lines, GetPxDisplayLines(generate_req.pxl_script()));
  if (display_lines.size() != schema.size()) {
    return error::InvalidArgument("script has $0 output tables, but $1 px.display calls",
                                  schema.size(), display_lines.size());
  }

  // Verify that table names are fully unique.
  absl::flat_hash_set<std::string> table_names;
  for (const auto& display_call : display_lines) {
    if (table_names.contains(display_call.table_name)) {
      return error::InvalidArgument("duplicate table name. '$0' already in use",
                                    display_call.table_name);
    }
    table_names.insert(display_call.table_name);
  }

  PL_ASSIGN_OR_RETURN(auto unique_df_name, GetUnusedVarName(generate_req.logical_planner_state(),
                                                            generate_req.pxl_script(), "otel_df"));

  std::vector<std::string> all_script_lines = absl::StrSplit(generate_req.pxl_script(), '\n');
  int64_t prev_idx = 0;
  std::vector<std::string> blocks;
  for (const auto& [i, display_call] : Enumerate(display_lines)) {
    std::vector<std::string> out_lines;
    for (int64_t j = prev_idx; j <= display_call.line_number_end; ++j) {
      out_lines.push_back(all_script_lines[j]);
    }
    prev_idx = display_call.line_number_end + 1;
    out_lines.push_back(absl::Substitute("\n$0 = $1", unique_df_name, display_call.table_argument));

    if (!schema.contains(display_call.table_name)) {
      return error::InvalidArgument(
          "no relation generated for $0, likely that the table name in the px.display call is "
          "duplicated",
          display_call.table_name);
    }

    auto relation = schema.at(display_call.table_name);
    PL_ASSIGN_OR_RETURN(std::string export_statement,
                        RelationToOTelExport(display_call.table_name, unique_df_name, relation));
    out_lines.push_back(export_statement);

    blocks.push_back(absl::StrJoin(out_lines, "\n"));
  }
  // Grab the content after the last display call and add it as the last block.
  std::vector<std::string> out_lines;
  for (int64_t j = prev_idx; j < static_cast<int64_t>(all_script_lines.size()); ++j) {
    out_lines.push_back(all_script_lines[j]);
  }
  if (out_lines.size() > 0) {
    blocks.push_back(absl::StrJoin(out_lines, "\n"));
  }

  auto generate_resp = std::make_unique<plannerpb::GenerateOTelScriptResponse>();
  generate_resp->set_otel_script(absl::StrJoin(blocks, "\n\n"));
  return generate_resp;
}

}  // namespace planner
}  // namespace carnot
}  // namespace px
