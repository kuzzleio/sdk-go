#!/usr/bin/env bash
set -e
cd internal/wrappers
./build_cpp_tests.sh
./_build_cpp_tests/KuzzleSDKStepDefs &

if [ ! -z "$FEATURE_FILE" ]; then
  cucumber features/$FEATURE_FILE
else
  cucumber
fi
cd -
