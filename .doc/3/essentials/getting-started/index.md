---
code: false
type: page
title: Getting started
description: Getting started
order: 99
---

# Getting Started

In this tutorial you will learn how to install the Kuzzle **Go SDK**.
This page shows examples of scripts that **store** documents in Kuzzle, and of scripts that subcribe to real-time **notifications** for each new document created.

::: success
Before proceeding, please make sure your system meets the following requirements:

- **Go** version 1.9 or higher ([Go installation instructions](https://golang.org/doc/install))
- A running Kuzzle server ([Kuzzle installation guide](core/2/guides/getting-started/run-kuzzle))
  :::

::: info
Having trouble? Get in touch with us on [Discord](http://join.discord.kuzzle.io)!
:::

## Installation

To easily install the Go SDK:

```bash
$ go get github.com/kuzzleio/sdk-go
```

This fetches the SDK and installs it in your `GOPATH` directory.

## First connection

Initialize a new Go project as described in the [Go Documentation](https://golang.org/doc/code.html#Command).
Then create a `init.go` file and start by adding the code below:

<<< ./snippets/init.go

This program initializes the Kuzzle server storage by creating a index, and a collection inside it
Run the program with the following command:

```bash
$ go run init.go
Connected!
Index nyc-open-data created!
Collection yellow-taxi created!
```

Congratulations, you performed a first connection to Kuzzle with a Go program.
You are now able to:

- Load the `Kuzzle Go SDK` from your `GOPATH` directory
- Instantiate a protocol (here `websocket`) and a Kuzzle SDK instance
- Connect to a Kuzzle instance running on `localhost`, with the WebSocket protocol
- Create a index
- Create a collection within an existing index

## Create your first document

Now that you successfully connected to your Kuzzle server with the Go SDK, and created an index and a collection, it's time to manipulate data.

Here is how Kuzzle structures its storage space:

- indexes contain collections
- collections contain documents
  Create a `document.go` file in the playground and add this code:

<<< ./snippets/document.go

As you did before, run your program:

```bash
$ go run document.go
Connected!
New document added to yellow-taxi collection!
```

You can perform other actions such as [delete](/sdk/go/3/controllers/document/delete),
[replace](/sdk/go/3/controllers/document/replace) or [search](/sdk/go/3/controllers/document/search) documents. There are also other ways to interact with Kuzzle like our [Admin Console](http://console.kuzzle.io), the [Kuzzle HTTP API](/core/2/api/protocols/http) or by using your [own protocol](/core/2/guides/write-protocols/start-writing-protocols).

Now you know how to:

- Store documents in a Kuzzle server, and access those

## Subscribe to realtime document notifications (pub/sub)

Time to use realtime with Kuzzle. Create a new file `realtime.go` with the following code:

<<< ./snippets/realtime.go

This program subscribes to changes made to documents with a `license` field set to `B`, within the `yellow-taxi` collection. Whenever a document matching the provided filters changes, a new notification is received from Kuzzle.
Run your program:

```bash
$ go run realtime.go
Connected!
Successfully subscribing!
New document added to yellow-taxi collection!
Driver John born on 1995-11-27 got a B license.
```

Now, you know how to:

- Create realtime filters
- Subscribe to notifications

## Where do we go from here?

Now that you're more familiar with the Go SDK, you can dive even deeper to learn how to leverage its full capabilities:

- discover what this SDK has to offer by browsing other sections of this documentation
- learn how to use [Koncorde](/core/2/api/koncorde-filters-syntax) to create incredibly fine-grained and blazing-fast subscriptions
- follow our guide to learn how to perform [basic authentication](/core/2/guides/main-concepts/authentication#local-strategy)
- follow our guide to learn how to [manage users and how to set up fine-grained access control](/core/2/guides/main-concepts/permissions)
