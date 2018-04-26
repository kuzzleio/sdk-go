[![Build Status](https://travis-ci.org/kuzzleio/sdk-go.svg?branch=master)](https://travis-ci.org/kuzzleio/sdk-go) [![codecov.io](http://codecov.io/github/kuzzleio/sdk-php/coverage.svg?branch=master)](http://codecov.io/github/kuzzleio/sdk-go?branch=master) [![GoDoc](https://godoc.org/github.com/kuzzleio/sdk-go?status.svg)](https://godoc.org/github.com/kuzzleio/sdk-go)

Official Kuzzle GO SDK with wrappers for C++ and JAVA SDK
======

## About Kuzzle

A backend software, self-hostable and ready to use to power modern apps.

You can access the Kuzzle repository on [Github](https://github.com/kuzzleio/kuzzle)

* [SDK Documentation](https://godoc.org/github.com/kuzzleio/sdk-go)
* [Installation](#installation)
* [Basic usage](#basic-usage)
* [Running tests](#tests)
* [License](#license)

## SDK Documentation

The complete SDK documentation is available [here](http://docs.kuzzle.io/sdk-reference/)

## Installation

````sh
go get github.com/kuzzleio/sdk-go
````

## Basic usage

````go
func main() {
    conn := websocket.NewWebSocket("localhost:7512", nil)
    k, _ := kuzzle.NewKuzzle(conn, nil)

    res, err := k.GetAllStatistics(nil)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, stat := range res {
        fmt.Println(stat.Timestamp, stat.FailedRequests, stat.Connections, stat.CompletedRequests, stat.OngoingRequests)
    }
}


````

## <a name="tests"></a> Running Tests

### Unit tests

To run the tests you can simply execute the coverage.sh script
```sh
./coverage.sh
```

You can also get html coverage by running
```sh
./coverage.sh --html
```
### e2e tests

To run e2e tests ensure you have a kuzzle running and then run
```sh
./internal/wrappers/features/e2e.sh
```

## Wrappers

### Dependencies

Before generating the wrappers you will need to install:

- [swig](www.swig.org)
- [Java 8](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) (don't forget to set your JAVA_HOME environment variable)

### Generate

## JAVA

```sh
make java
```

You will find the final jars files in `internal/wrappers/build/java/build/libs`

## CP

```sh
make cpp
```
You will find the final .so file in `internal/wrappers/build/cpp`

## All at once

You can generate all sdk's at once by typing

```sh
make all
```

You will be able to find the final 

## License

[Apache 2](LICENSE.md)
