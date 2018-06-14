#!/usr/bin/env bash
set -ex
# Download and launch custom Kuzzle stack
KUZZLE_CHECK_CONNECTIVITY_CMD="curl -o /dev/null http://localhost:7512"

docker-compose -f .codepipeline/docker-compose.yml up -d

while ! $KUZZLE_CHECK_CONNECTIVITY_CMD &> /dev/null
  do
    sleep 2
  done
