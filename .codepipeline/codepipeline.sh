#!/usr/bin/env bash
set -e
# Build all languages
echo -e "\n\n        [ Building & Testing ]\n"

docker run --rm --privileged --name build-machine \
             --mount type=bind,source="$(pwd)",target=/go/src/github.com/kuzzleio/sdk-go \
             kuzzleio/sdk-cross:$1 /build.sh


docker run --rm --network codepipeline_default --link kuzzle --name build-machine \
          --mount type=bind,source="$(pwd)",target=/go/src/github.com/kuzzleio/sdk-go \
          kuzzleio/sdk-cross:$1 /test.sh && \

docker kill codepipeline_elasticsearch_1 codepipeline_redis_1 kuzzle
