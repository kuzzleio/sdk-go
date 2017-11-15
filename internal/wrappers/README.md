[![Build Status](https://travis-ci.org/kuzzleio/sdk-go.svg?branch=master)](https://travis-ci.org/kuzzleio/sdk-go)

# SDK Wrappers

This project contains a CGO wrapper to Kuzzle's [Go SDK](https://github.com/kuzzleio/sdk-go) and [SWIG templates](http://www.swig.org/), allowing to publish Kuzzle's SDK in multiple languages.

# Pre-requisites

* [Go](https://golang.org/doc/install)
* [JSON-C library](https://github.com/json-c/json-c#install-using-apt-eg-ubuntu-16042-lts)

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
