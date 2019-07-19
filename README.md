<p align="center">
  <img src="https://user-images.githubusercontent.com/7868838/53850936-31e57180-3fbd-11e9-8392-8f3e26bf2aa8.png"/>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/tested%20on-linux%20%7C%20osx%20%7C%20windows-blue.svg">
  <a href="https://travis-ci.org/kuzzleio/sdk-go">
    <img src="https://travis-ci.org/kuzzleio/sdk-go.svg?branch=master"/>
  </a>
  <a href="https://codecov.io/gh/kuzzleio/sdk-go">
    <img src="https://codecov.io/gh/kuzzleio/sdk-go/branch/master/graph/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/kuzzleio/sdk-go">
    <img src="https://goreportcard.com/badge/github.com/kuzzleio/sdk-go" />
  </a>
  <a href="https://godoc.org/github.com/kuzzleio/sdk-go">
    <img src="https://godoc.org/github.com/kuzzleio/sdk-go?status.svg"/>
  </a>
  <a href="https://github.com/kuzzleio/sdk-go/blob/master/LICENSE">
    <img alt="undefined" src="https://img.shields.io/github/license/kuzzleio/sdk-go.svg?style=flat">
  </a>
</p>

## About

### Kuzzle Go

This is the official Go SDK for the free and open-source backend Kuzzle. It provides a way to dial with a Kuzzle server from Go applications.
The SDK provides a native __WebSocket__ support. You can add your own network protocol by implementing the Protocol interface.

<p align="center">
  :books: <b><a href="https://docs.kuzzle.io/sdk-reference/go/1">Documentation</a></b>
</p>

### Kuzzle

Kuzzle is a ready-to-use, **on-premises and scalable backend** that enables you to manage your persistent data and be notified in real-time on whatever happens to it. 
It also provides you with a flexible and powerful user-management system.

* :watch: __[Kuzzle in 5 minutes](https://kuzzle.io/company/about-us/kuzzle-in-5-minutes/)__
* :octocat: __[Github](https://github.com/kuzzleio/kuzzle)__
* :earth_africa: __[Website](https://kuzzle.io)__
* :books: __[Documentation](https://docs.kuzzle.io)__
* :email: __[Gitter](https://gitter.im/kuzzleio/kuzzle)__

### Get trained by the creators of Kuzzle :zap:

Train yourself and your teams to use Kuzzle to maximize its potential and accelerate the development of your projects.  
Our teams will be able to meet your needs in terms of expertise and multi-technology support for IoT, mobile/web, backend/frontend, devops.  
:point_right: [Get a quote](https://hubs.ly/H0jkfJ_0)

## Usage

### Installation

Simply download the SDK to your `GOPATH`.

```go
go get github.com/kuzzleio/sdk-go
```

### Example

The SDK supports different protocols. When instantiating, 
you must choose the protocol to use and fill in the different options needed to connect to Kuzzle.  

```go
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

	timestamp, err := k.Server.Now(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(timestamp)
}
```

## Contributing

First of all, thank you to take the time to contribute to this SDK. To help us validating your future pull request,
please make sure your work pass linting and unit tests.

```bash 
$ bash .ci/test_with_coverage.sh 
```

If you want to see current coverage run the script with this argument.

```bash 
$ bash .ci/test_with_coverage.sh  --html
```

This should open a new tab in your favorite web browser and allow you to see the lines of code covered by the unit tests.


