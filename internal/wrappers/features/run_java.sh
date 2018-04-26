#!/bin/bash

cd internal/wrappers/features/java
taskset -c 1 gradle --stacktrace cucumber
cd -
