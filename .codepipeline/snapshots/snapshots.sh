#!/usr/bin/env bash

set -e

dirs=(java cpp c)

function package_and_push_java() {
  sdk_version=$1

  cd internal/wrappers/build/java/build/libs

  sdk_name="kuzzlesdk-java.jar"
  s3_dest="s3://$AWS_S3_BUCKET/sdk/java/$sdk_version/$sdk_name"

  mv kuzzlesdk-1.0.0.jar $sdk_name

  aws s3 cp $sdk_name $s3_dest

  cd -
}

function package_and_push_c() {
  sdk_version=$1

  cd internal/wrappers/build/c

  sdk_name="libkuzzlesdk-c.so"
  s3_dest="s3://$AWS_S3_BUCKET/sdk/c/$sdk_version/kuzzlesdk-c.tar.gz"

  mv libkuzzlesdk.so $sdk_name
  mkdir lib include
  cp ../../headers/*.h include
  cp $sdk_name lib/
  tar cfz kuzzlesdk-c.tar.gz lib include

  aws s3 cp kuzzlesdk-c.tar.gz $s3_dest

  cd -
}

function package_and_push_cpp() {
  sdk_version=$1

  cd internal/wrappers/build/cpp

  sdk_name="libkuzzlesdk.so"
  s3_dest="s3://$AWS_S3_BUCKET/sdk/cpp/$sdk_version/kuzzlesdk-cpp.tar.gz"

  mv libcpp.so $sdk_name
  mkdir lib include
  cp ../../headers/* include
  cp ../c/kuzzle.h include/kuzzle.h
  cp $sdk_name lib/
  tar cfz kuzzlesdk-cpp.tar.gz lib include

  aws s3 cp kuzzlesdk-cpp.tar.gz $s3_dest

  cd -
}

function package_and_push_sdks() {
  sdk_version=$1

  for dir in ${dirs[@]}; do
    echo -e "\n----------------------------------------------------------------\n"
    figlet "$dir SDK Snapshot"
    echo -e "\n----------------------------------------------------------------\n"


    package_and_push_$dir $sdk_version

    if [[ $? -ne 0 ]]; then
      exit 1
    fi
  done

  if [[ $? -eq 0 ]]; then
    aws cloudfront create-invalidation --distribution-id $AWS_CLOUDFRONT_DISTRIBUTION_ID --paths "/*"
  fi
}

if [[ $TRAVIS_PULL_REQUEST = false ]]; then
  echo "This script run only on push and not on PR"
  exit 0
fi

if [[ $TRAVIS_BRANCH = "master" ]]; then
  sdk_version="latest"
  package_and_push_sdks $sdk_version
elif [[ $TRAVIS_BRANCH = *"-stable" ]]; then
  sdk_version=$TRAVIS_BRANCH
  package_and_push_sdks $sdk_version
elif [[ $TRAVIS_BRANCH = *"-dev" ]]; then
  sdk_version=$TRAVIS_BRANCH
  package_and_push_sdks $sdk_version
else
  echo "TRAVIS_BRANCH not set or not on a build branch (*-stable, *-dev or master)"
fi
