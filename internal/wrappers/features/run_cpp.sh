#!/usr/bin/env bash
set -e
cd internal/wrappers
./build_cpp_tests.sh̀
./_build_cpp_tests/KuzzleSDKStepDefs > /dev/null &&
cucumber
cd -
