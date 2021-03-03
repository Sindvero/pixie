#!/bin/bash
set -e

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
uipath="${SCRIPTPATH}/../src/ui"

pushd "${uipath}" > /dev/null
./node_modules/.bin/eslint "$@"
popd > /dev/null
