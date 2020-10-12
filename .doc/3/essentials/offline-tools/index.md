---
code: false
type: page
title: Offline Tools
description: Tools to handle the loss of connection to the Kuzzle server
order: 100
---

# Offline tools

The Kuzzle SDK provides a set of properties that help your application to be resilient to the loss of network connection
during its lifespan.

## offlineQueue

A read-only list of `QueryObject` containing the requests queued while the SDK is in the `offline` state (it behaves like a FIFO queue).

## queueMaxSize

A writable `int` defining the maximun size of the `offlineQueue`.

## queueTTL

A writable `Duration` defining the time in milliseconds a queued request is kept in the `offlineQueue`.

## StartQueuing()

Starts the requests queuing. Request will be put in the `offlineQueue` instead of being discarded, until `stopQueuing` is called.
Works only in `offline` state, and if the `autoQueue` option is set to false. Call `playQueue` to send to Kuzzle the
requests in the queue, once the SDK state passes to `online`. Call `flushQueue` to empty the queue without sendint the requests.

## StopQueuing()

Stop queuing the requests. Requests will no more be put in the `offlineQueue`, they will be discarded.
Works only in the `offline` state, and if the `autoQueue` option is set to `false`.

## PlayQueue()

Sends to Kuzzle all the requests in the `offlineQueue`. Works only if the SDK is not in a `offline` state, and if the
`autoReplay` option is set to false.

## FlushQueue()

Empties the `offlineQueue` without sending the requests to Kuzzle.

## autoQueue

A writable `bool` telling the SDK whether to automatically queue requests during the `offline` state or not.

## autoReplay

A writable `bool` telling the SDK whether to automatically send or not the requests in the `offlineQueue` on a
`reconnected` event.

## autoReconnect

A writable `bool` telling the SDK whether to automatically reconnect or not to Kuzzle after a connection loss.

## reconnectionDelay

A read-only `Duration` specifying the time in milliseconds between different reconnection attempts.

## autoResubscribe

A writable `bool` telling the SDK whether to automatically renew or not all subscriptions on a reconnected event.

## queueFilter

A writable `QueueFilter` called by the SDK each time a `Request` need to be queued. The `Request` is passed as the only argument
to the function and is queued only if the function returns `true`. Use it to define which requests are allowed to be queued.

## offlineQueueLoader

A writable `OfflineQueueLoader` called by the SDK before playing the requests in the `offlineQueue`. This function takes no arguments
and returns an array of `KuzzleRequest` that are added on top of the `offlineQueue`. Use it to inject new requests to be played
before the queue.
