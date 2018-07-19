#!/usr/bin/env bash
set -e
cd internal/wrappers
sh ./build_cpp_tests.sh

FEATURE_FILE=$1

if [ -z "$KUZZLE_HOST" ]; then
  # For debug
  ./_build_cpp_tests/KuzzleSDKStepDefs &
else
  # To hide cucumber debug output in CI
  ./_build_cpp_tests/KuzzleSDKStepDefs > /dev/null &
fi

if [ ! -z "$FEATURE_FILE" ]; then
  cucumber features/$FEATURE_FILE
else
  cucumber
fi

cd -
