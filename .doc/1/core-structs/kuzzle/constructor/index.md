---
code: true
type: page
title: Constructor
description: Creates a new Kuzzle object connected to the backend
order: 100
---

# Constructor

This is the main entry point to communicate with Kuzzle.
Each instance represents a connection to Kuzzle with specific options.

This interface implements the [KuzzleEventEmitter](/sdk/go/1/core-structs/kuzzle-event-emitter) interface

## Arguments

```go
NewKuzzle(protocol connection.Connection) (*Kuzzle, error)
```

| Argument   | Type                  | Description                           |
| ---------- | --------------------- | ------------------------------------- |
| `protocol` | <pre>connection.Connection</pre> | The protocol used by the SDK instance |

### **protocol**

A [Protocol](/sdk/go/1/protocols/) is a structure implementing the `connection.Connection` interface.
The available protocols are:

- `websocket.Websocket`

The protocol must be instantiated and passed to the constructor.
It takes the following arguments:

| Argument  | Type          | Description                     | Required |
| --------- | ------------- | ------------------------------- | -------- |
| `host`    | <pre>string</pre>        | Kuzzle hostname to connect to   | yes      |
| `options` | <pre>types.Options</pre> | Kuzzle connection configuration | yes      |

The `options` parameter of the protocol constructor has the following properties.
You can use standard getter/setter to use these properties.

| Option              | Type         | Description                                                        | Default        | Required |
| ------------------- | ------------ | ------------------------------------------------------------------ | -------------- | -------- |
| `autoQueue`         | <pre>bool</pre>         | Automatically queue all requests during offline mode               | `false`        | no       |
| `autoReconnect`     | <pre>bool</pre>         | Automatically reconnect after a connection loss                    | `true`         | no       |
| `autoReplay`        | <pre>bool</pre>         | Automatically replay queued requests on a `reconnected` event      | `false`        | no       |
| `autoResubscribe`   | <pre>bool</pre>         | Automatically renew all subscriptions on a `reconnected` event     | `true`         | no       |
| `offlineMode`       | <pre>int</pre>          | Offline mode configuration. `types.Manual` or `types.Auto`         | `types.Manual` | no       |
| `port`              | <pre>int</pre>          | Target Kuzzle port                                                 | `7512`         | no       |
| `queueTTL`          | <pre>int</pre>          | Time a queued request is kept during offline mode, in milliseconds | `120000`       | no       |
| `queueMaxSize`      | <pre>int</pre>          | Number of maximum requests kept during offline mode                | `500`          | no       |
| `replayInterval`    | <pre>Duration</pre>     | Delay between each replayed requests, in milliseconds              | `10`           | no       |
| `reconnectionDelay` | <pre>Duration</pre>     | number of milliseconds between reconnection attempts               | `1000`         | no       |
| `sslConnection`     | <pre>bool</pre>         | Switch Kuzzle connection to SSL mode                               | `false`        | no       |
| `volatile`          | <pre>VolatileData</pre> | Common volatile data, will be sent to all future requests          | -              | no       |

## Getter & Setter

These properties of the Kuzzle struct can be writable.
For example, you can read the `volatile` property via `getVolatile()` and set it via `setVolatile()`.

| Property name        | Type               | Description                                                                                                               | Availability |
| -------------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------- | :----------: |
| `autoQueue`          | <pre>bool</pre>               | Automatically queue all requests during offline mode                                                                      |   Get/Set    |
| `autoReconnect`      | <pre>bool</pre>               | Automatically reconnect after a connection loss                                                                           |     Get      |
| `autoReplay`         | <pre>bool</pre>               | Automatically replay queued requests on a `reconnected` event                                                             |   Get/Set    |
| `autoResubscribe`    | <pre>bool</pre>               | Automatically renew all subscriptions on a `reconnected` event                                                            |   Get/Set    |
| `host`               | <pre>string</pre>             | Target Kuzzle host                                                                                                        |     Get      |
| `port`               | <pre>int</pre>                | Target Kuzzle port                                                                                                        |     Get      |
| `jwt`                | <pre>string</pre>             | Token used in requests for authentication.                                                                                |   Get/Set    |
| `offlineQueue`       | <pre>QueryObject</pre>        | Contains the queued requests during offline mode                                                                          |     Get      |
| `offlineQueueLoader` | <pre>OfflineQueueLoader</pre> | Called before dequeuing requests after exiting offline mode, to add items at the beginning of the offline queue           |   Get/Set    |
| `queueFilter`        | <pre>QueueFilter</pre>        | Called during offline mode. Takes a request object as arguments and returns a bool, indicating if a request can be queued |   Get/Set    |
| `queueMaxSize`       | <pre>int</pre>                | Number of maximum requests kept during offline mode                                                                       |   Get/Set    |
| `queueTTL`           | <pre>Duration</pre>           | Time a queued request is kept during offline mode, in milliseconds                                                        |   Get/Set    |
| `replayInterval`     | <pre>Duration</pre>           | Delay between each replayed requests                                                                                      |   Get/Set    |
| `reconnectionDelay`  | <pre>Duration</pre>           | Number of milliseconds between reconnection attempts                                                                      |     Get      |
| `sslConnection`      | <pre>bool</pre>               | Connect to Kuzzle using SSL                                                                                               |     Get      |
| `volatile`           | <pre>VolatileData</pre>       | Common volatile data, will be sent to all future requests                                                                 |   Get/Set    |

**Notes:**

- multiple methods allow passing specific `volatile` data. These `volatile` data will be merged with the global Kuzzle `volatile` object when sending the request, with the request specific `volatile` taking priority over the global ones.
- the `queueFilter` property is a function taking a `QueryObject` as an argument. This object is the request sent to Kuzzle, following the [Kuzzle API](/core/1/api/essentials/query-syntax) format
- if `queueTTL` is set to `0`, requests are kept indefinitely
- The offline buffer acts like a first-in first-out (FIFO) queue, meaning that if the `queueMaxSize` limit is reached, older requests are discarded to make room for new requests
- if `queueMaxSize` is set to `0`, an unlimited number of requests is kept until the buffer is flushed
- the `offlineQueueLoader` must be set with a function, taking no argument, and returning an array of objects containing a `query` member with a Kuzzle query to be replayed, and an optional `cb` member with the corresponding callback to invoke with the query result
- updates to `autoReconnect`, `reconnectionDelay` and `sslConnection` properties will only take effect on next `connect` call

## Return

A `Kuzzle` struct and an [error struct](/sdk/go/1/essentials/error-handling).
The `error` struct is nil if everything was ok.

## Usage

In a first step, you have to create a new `connection.Connection` and pass it to the constructor.
By now the only connection available is `websocket.Websocket`.

<<< ./snippets/constructor.go
