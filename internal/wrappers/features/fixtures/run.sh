#!/bin/bash
set -e
CONTENT_TYPE="Content-Type: application/json"

KUZZLE_HOST=http://localhost:7512

# update rights for anonymous user
curl -X PUT -H "$CONTENT_TYPE" -d @rights.json $KUZZLE_HOST/roles/anonymous/_update

# create index
curl -X POST -H "$CONTENT_TYPE" $KUZZLE_HOST/index/_create

# create collection
curl -X PUT -H "$CONTENT_TYPE" $KUZZLE_HOST/index/collection

# create geofence collection
curl -X PUT -H "$CONTENT_TYPE" $KUZZLE_HOST/index/geofence

#update mapping for geofence collection
curl -X PUT -H "$CONTENT_TYPE" -d @mapping.json $KUZZLE_HOST/index/geofence/_mapping
