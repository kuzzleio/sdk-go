#!/usr/bin/env bash
set -e
cd internal/wrappers
./build_cpp_tests.sh

FEATURE_FILE=$1

if [ ! -z "$KUZZLE_HOST" ]; then
  ./_build_cpp_tests/KuzzleSDKStepDefs &
else
  ./_build_cpp_tests/KuzzleSDKStepDefs > /dev/null &
fi

if [ ! -z "$FEATURE_FILE" ]; then
  cucumber features/$FEATURE_FILE
else
  cucumber
fi

cd -
