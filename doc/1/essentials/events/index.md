---
code: false
type: page
title: Events
description: SDK events system
order: 100
---

# Events

An event system allows to be notified when the SDK status changes. These events are issued by the [Kuzzle](/sdk/go/1/core-structs/kuzzle) interface.

The API for interacting with events is described by our [KuzzleEventEmitter](/sdk/go/1/core-structs/kuzzle-event-emitter) interface documentation.

# Emitted Events

The following event identifiers are constants declared in the `event` package.

## Connected

Triggered when the SDK has successfully connected to Kuzzle.

## Discarded

Triggered when Kuzzle rejects a request (e.g. request can't be parsed, request too large, ...).

**Channel signature:** `chan<- *types.KuzzleResponse)`

## Disconnected

Triggered when the current session has been unexpectedly disconnected.

**Channel signature:** `chan<- interface{}` (will receive nil)

## LoginAttempt

Triggered when a login attempt completes, either with a success or a failure result.

**Channel signature:** `chan<- *types.LoginAttempt`

## NetworkError

Triggered when the SDK has failed to connect to Kuzzle.
This event does not trigger the offline mode.

**Channel signature:** `chan<- error`

## OfflineQueuePop

Triggered whenever a request is removed from the offline queue.

**Channel signature:** `chan<- *types.QueryObject`

## OfflineQueuePush

Triggered whenever a request is added to the offline queue.

**Channel signature:** `chan<- *types.QueryObject`

## QueryError

Triggered whenever Kuzzle responds with an error

**Channel signature:** `chan<- *types.QueryObject`

## Reconnected

Triggered when the current session has reconnected to Kuzzle after a disconnection, and only if ``AutoReconnect`` is set to ``true``.

**Channel signature:** `chan<- interface{}` (will receive nil)

## TokenExpired

Triggered when Kuzzle rejects a request because the authentication token has expired.

**Channel signature:** `chan<- interface{}` (will receive nil)
