#!/bin/bash

CONTENT_TYPE="Content-Type: application/json"
KUZZLE_HOST=http://localhost:7512

# update rights for anonymous user
curl -X PUT -H $CONTENT_TYPE -d "`cat rights.json`" $KUZZLE_HOST/roles/anonymous/_update

# create index
curl -X POST -H $CONTENT_TYPE $KUZZLE_HOST/index/_create

# create collection
curl -X PUT -H $CONTENT_TYPE $KUZZLE_HOST/index/collection
