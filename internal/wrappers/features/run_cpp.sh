#!/usr/bin/env bash
set -e
cd internal/wrappers
sh ./build_cpp_tests.sh

FEATURE_FILE=$1

./_build_cpp_tests/KuzzleSDKStepDefs > /dev/null &

if [ ! -z "$FEATURE_FILE" ]; then
  cucumber features/$FEATURE_FILE
else
  cucumber
fi

cd -
