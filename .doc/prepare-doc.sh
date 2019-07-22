#!/bin/bash

echo "Clone documentation framework"
git clone --depth 10 --single-branch --branch master https://github.com/kuzzleio/documentation.git framework/

echo "Link local documentation"
rm framework/src/sdk/go/1 # remove link to submodule
ln -s ../../../../1 framework/src/sdk/go/1 # use current documentation

npm --prefix framework/ install