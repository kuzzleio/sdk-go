[![Build Status](https://travis-ci.org/kuzzleio/sdk-go.svg?branch=master)](https://travis-ci.org/kuzzleio/sdk-go) [![codecov.io](http://codecov.io/github/kuzzleio/sdk-php/coverage.svg?branch=master)](http://codecov.io/github/kuzzleio/sdk-go?branch=master) [![GoDoc](https://godoc.org/github.com/kuzzleio/sdk-go?status.svg)](https://godoc.org/github.com/kuzzleio/sdk-go)

Official Kuzzle GO SDK
======

## About Kuzzle

A backend software, self-hostable and ready to use to power modern apps.

You can access the Kuzzle repository on [Github](https://github.com/kuzzleio/kuzzle)

* [SDK Documentation](https://godoc.org/github.com/kuzzleio/sdk-go)
* [Report an issue](#report-an-issue)
* [Installation](#installation)
* [Basic usage](#basic-usage)
* [Running tests](#tests)
* [License](#license)

## SDK Documentation

The complete SDK documentation is available [here](http://docs.kuzzle.io/sdk-reference/)

## Report an issue

Use following meta repository to report issues on SDK:

https://github.com/kuzzleio/kuzzle-sdk/issues

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

To run the tests you can simply execute the coverage.sh script
```sh
./coverage.sh
```

You can also get html coverage by running
```sh
./coverage.sh --html
```

## License

[Apache 2](LICENSE.md)
