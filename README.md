[![Build Status](https://travis-ci.org/kuzzleio/sdk-go.svg?branch=master)](https://travis-ci.org/kuzzleio/sdk-go) [![codecov.io](http://codecov.io/github/kuzzleio/sdk-php/coverage.svg?branch=master)](http://codecov.io/github/kuzzleio/sdk-go?branch=master) [![GoDoc](https://godoc.org/github.com/kuzzleio/sdk-go?status.svg)](https://godoc.org/github.com/kuzzleio/sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kuzzleio/sdk-go)](https://goreportcard.com/report/github.com/kuzzleio/sdk-go)

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
./test.sh
```

You can also get html coverage by running
```sh
./test.sh --html
```
### e2e tests

#### JAVA

```sh
cd internal/wrappers
make java
cd features/java
gradle cucumber
```

#### C++

```sh
make cpp
cd internal/wrappers
./build_cpp_tests.sh̀
./_build_cpp_tests/KuzzleSDKStepDefs > /dev/null &
cucumber
```



## Wrappers

### Dependencies

Before generating the wrappers you will need to install:

- You will need a g++ compatible C++11
- [swig](www.swig.org)
- [Java 8](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) (don't forget to set your JAVA_HOME environment variable)
- Python You will need to install python-dev to compile the python SDK

### Generate

## JAVA

```sh
make java
```

You will find the final jars files in `internal/wrappers/build/java/build/libs`

## CPP

```sh
make cpp
```
You will find the final .so file in `internal/wrappers/build/cpp`

## Python
```sh
make python
```
You will find the final .so file in `internal/wrappers/build/python`

## CSHARP

### Build on Windows

#### Prerequisites
- Visual Studio 2017
- Windows SDK
- Go - https://golang.org/doc/install
- Mono (x64) - https://www.mono-project.com/download/stable/
- Make (GNU - Windows) - http://gnuwin32.sourceforge.net/packages/make.htm
- MinGW (x64 - Posix) - https://sourceforge.net/projects/mingw-w64/files/mingw-w64/mingw-w64-release/

#### Compiling Csharp
- Add Go/Mono/Make/MinGW intallation directory to PATH
- Run Visual Studio developper command line tool
- Execute `make csharp`

## All at once

You can generate all sdk's at once by typing

```sh
make all
```

## Generate wrappers and launch e2e tests using Docker

You can use Docker to simplify wrappers generation and testing

### Build

In project root, use:

```bash
$ docker run --rm -it -v "$(pwd)":/go/src/github.com/kuzzleio/sdk-go kuzzleio/sdk-cross:amd64 /build.sh
```

This command will build all wrappers using our Docker Image

### Testing

E2E tests need a running Kuzzle instance so run the script:

```bash
$ sh .codepipeline/start_kuzzle.sh
```

Now run tests using Docker:

```bash
$ docker run --rm -it --network codepipeline_default --link kuzzle -v "$(pwd)":/go/src/github.com/kuzzleio/sdk-go kuzzleio/sdk-cross:amd64 /test.sh
```

## License

[Apache 2](LICENSE.md)
