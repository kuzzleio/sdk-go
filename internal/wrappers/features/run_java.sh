#!/bin/bash
set -e
cd internal/wrappers/features/java
gradle cucumber
cd -
