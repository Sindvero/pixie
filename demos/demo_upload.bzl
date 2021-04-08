def _impl(ctx):
    bucket = ctx.attr.bucket
    if bucket[-1] == '/':
        fail('Bucket name must not end with "/"')

    manifest_cmd = """
   	gsutil -h "Cache-Control:no-cache,max-age=0" \
           -h "Content-Type:application/json" \
           cp {} {}/manifest.json
    """.format(ctx.file.manifest.short_path, ctx.attr.bucket)

    archive_cmds = []
    for archive in ctx.files.archives:
        archive_cmds.append("""
        gsutil -h "Cache-Control:no-cache,max-age=0" \
               -h "Content-Type:application/gzip" \
               cp {} {}/{}
        """.format(archive.short_path, ctx.attr.bucket, archive.basename))

    cmds = ["#!/bin/sh -e\n", manifest_cmd] + archive_cmds
    ctx.actions.write
    ctx.actions.write(
        output=ctx.outputs.executable,
        content="\n".join(cmds),
    )
    runfiles = ctx.runfiles(files=[ctx.file.manifest] + ctx.files.archives)
    return [DefaultInfo(runfiles=runfiles)]


_demo_upload = rule(
    attrs={
        'bucket': attr.string(
            mandatory=True,
            doc='Target GCS bucket (string), e.g. gs://foo/bar',
        ),
        'manifest': attr.label(
            allow_single_file=True,
            mandatory=True,
            doc='The JSON manifest file for the demo application.',
        ),
        'archives': attr.label_list(
            doc='The tar files for the actual demos.',
        ),
    },
    executable=True,
    implementation=_impl,
    doc='Upload demos to GCS.',
)


def demo_upload(name, bucket, manifest, archives):
    _demo_upload(
        name=name,
        bucket=bucket,
        manifest=manifest,
        archives=archives,
    )