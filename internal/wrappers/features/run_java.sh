#!/bin/bash

cd internal/wrappers/features/fixtures/
./run.sh
cd ../java
taskset -c 1 gradle cucumber