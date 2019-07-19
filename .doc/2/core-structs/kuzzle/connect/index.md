---
code: true
type: page
title: connect
description: Connects the SDK to Kuzzle
---

# Connect

Connects to Kuzzle using the `host` argument provided to the `connection.Connection` (see [Kuzzle constructor](/sdk/go/1/core-structs/kuzzle/constructor/#usage-go)).
Subsequent call have no effect if the SDK is already connected.

## Arguments

```go
Connect() error
```

## Return

Return a [Kuzzle error](/sdk/go/1/essentials/error-handling) if the SDK can not connect to Kuzzle.

## Usage

<<< ./snippets/connect.go
