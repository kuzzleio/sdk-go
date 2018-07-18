#!/usr/bin/env bash

set -e

dirs=(java cpp c)


function snap_java() {
  cd internal/wrappers/build/java/build/libs
  newname="kuzzlesdk-java.jar"
  s3_dest="s3://$AWS_S3_BUCKET/sdk/$dest_dir$pr_name/$newname"
  mv kuzzlesdk-1.0.0.jar $newname

  if [[ $TRAVIS_PULL_REQUEST -ne false ]]; then
    aws s3 cp $newname $s3_dest --expires "$(date -d '+2 weeks' --utc +'%Y-%m-%dT%H:%M:%SZ')"
  else
    aws s3 cp $newname $s3_dest
  fi

  cd -
}

function snap_c() {
  cd internal/wrappers/build/c
  newname="libkuzzlesdk-c.so"
  s3_dest="s3://$AWS_S3_BUCKET/sdk/$dest_dir$pr_name/kuzzlesdk-c.tar.gz"
  mv libkuzzlesdk.so $newname
  mkdir lib include
  cp ../../headers/*.h include
  cp $newname lib/
  tar cfz kuzzlesdk-c.tar.gz lib include

  if [[ $TRAVIS_PULL_REQUEST -ne false ]]; then
    aws s3 cp kuzzlesdk-c.tar.gz $s3_dest --expires "$(date -d '+2 weeks' --utc +'%Y-%m-%dT%H:%M:%SZ')"
  else
    aws s3 cp kuzzlesdk-c.tar.gz $s3_dest
  fi

  cd -
}

function snap_cpp() {
  cd internal/wrappers/build/cpp
  newname="libkuzzlesdk.so"
  s3_dest="s3://$AWS_S3_BUCKET/sdk/$dest_dir$pr_name/kuzzlesdk-cpp.tar.gz"
  mv libcpp.so $newname
  mkdir lib include
  cp ../../headers/* include
  cp $newname lib/
  tar cfz kuzzlesdk-cpp.tar.gz lib include

  if [[ $TRAVIS_PULL_REQUEST -ne false ]]; then
    aws s3 cp kuzzlesdk-cpp.tar.gz $s3_dest --expires "$(date -d '+2 weeks' --utc +'%Y-%m-%dT%H:%M:%SZ')"
  else
    aws s3 cp kuzzlesdk-cpp.tar.gz $s3_dest
  fi

  cd -
}


if [[ -z $TRAVIS_PULL_REQUEST ]]; then
  export dest_dir="nightly"
else
  pr_num="$(echo $TRAVIS_PULL_REQUEST_BRANCH | cut -d- -f2)"
  dir_name="KZL-$pr_num"
  export pr_name="/$dir_name"
  export dest_dir="snapshots"
fi

for dir in ${dirs[@]}; do
  echo -e "\n----------------------------------------------------------------\n"
  figlet "$dir SDK Snapshot"
  echo -e "\n----------------------------------------------------------------\n"


  snap_$dir

  if [[ $? -ne 0 ]]; then
    exit 1
  fi
done

if [[ $? -eq 0 ]]; then
  aws cloudfront create-invalidation --distribution-id $AWS_CLOUDFRONT_DISTRIBUTION_ID --paths "/*"
fi
