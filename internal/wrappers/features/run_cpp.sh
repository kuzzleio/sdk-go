#!/usr/bin/env bash
set -e
cd internal/wrappers
./build_cpp_tests.sh
./_build_cpp_tests/KuzzleSDKStepDefs &
cucumber
cd -
