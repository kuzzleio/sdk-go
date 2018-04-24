#!/usr/bin/env bash

# Build all languages
echo -e "\n\n        [ Building ]\n"
docker run --rm --name build-machine \
           --mount type=bind,source="$(pwd)",target=/go/src/github.com/kuzzleio/sdk-go \
           kuzzleio/sdk-cross:$1 /build.sh

# Test all languages
echo -e "\n\n        [ Testing ]\n"
docker run --rm --network codepipeline_default --link kuzzle --name build-machine \
          --mount type=bind,source="$(pwd)",target=/go/src/github.com/kuzzleio/sdk-go \
          kuzzleio/sdk-cross:$1 /test.sh

# Kill Kuzzle stack before exit
docker kill codepipeline_elasticsearch_1 codepipeline_redis_1 kuzzle
