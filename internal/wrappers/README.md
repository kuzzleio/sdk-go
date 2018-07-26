[![Build Status](https://travis-ci.org/kuzzleio/sdk-go.svg?branch=master)](https://travis-ci.org/kuzzleio/sdk-go)

# SDK Wrappers

This project contains a CGO wrapper to Kuzzle's [Go SDK](https://github.com/kuzzleio/sdk-go) and [SWIG templates](http://www.swig.org/), allowing to publish Kuzzle's SDK in multiple languages.

# Pre-requisites

* [Go](https://golang.org/doc/install)
* G++ 6.x

## C++ SDK

You can use Docker to run the C++ functional tests:
 - Build the SDK : `docker run --rm -it -v "$(pwd)":/go/src/github.com/kuzzleio/sdk-go kuzzleio/sdk-cross:amd64 /build.sh`
 - Run a Kuzzle stack : `sh .codepipeline/start_kuzzle.sh`
 - Build and run features : `docker run --rm -it --network codepipeline_default --link kuzzle -e KUZZLE_HOST=kuzzle -v "$(pwd)":/go/src/github.com/kuzzleio/sdk-go  kuzzleio/sdk-cross:amd64 sh internal/wrappers/features/run_cpp.sh`
`

You can specify a single feature file to be run by passing it as argument to the `run_***.sh` script.  
Run only the collection features : `docker run --rm -it --network codepipeline_default --link kuzzle -e KUZZLE_HOST=kuzzle -v "$(pwd)":/go/src/github.com/kuzzleio/sdk-go  kuzzleio/sdk-cross:amd64 sh internal/wrappers/features/run_cpp.sh collection.feature`

# Contributing

* Clone this project:

```sh
$ git clone git@github.com:kuzzleio/sdk-wrappers.git
```

* Build all SDKs:

```sh
$ make
```

* If you need to build the CGO wrapper only:

```sh
$ make core
```
