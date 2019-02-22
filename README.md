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

The complete SDK documentation is available [here](https://docs.kuzzle.io/sdk-reference)

## Installation

````sh
go get github.com/kuzzleio/sdk-go
````

## Basic usage

````go
package main

import (
	"fmt"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
)

func main() {
	conn := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)
	k.Connect()

	res, err := k.Server.GetAllStats(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(res))
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

## License

[Apache 2](LICENSE.md)
