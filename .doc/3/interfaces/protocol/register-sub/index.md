---
code: true
type: page
title: RegisterSub
description: Used when subscribing to store a listener.
---

# RegisterSub

Attaches a notifications listener to an existing subscription.

## Signature

```go
RegisterSub(channel string, roomId string, filters json.RawMessage, subscribeToSelf bool, notificationResult chan<- types.NotificationResult, onReconnectChannel chan<- interface{})
```

## Arguments

| Argument            | Type                              | Description                                         |
| ------------------- | --------------------------------- | --------------------------------------------------- |
| `channel`           | <pre>string</pre>     | Subscription channel identifier                     |
| `roomId`           | <pre>const std::string&</pre>     | Subscription room identifier                        |
| `filters`           | <pre>const std::string&</pre>     | Subscription filters                                |
| `subscribeToSelf` | <pre>bool</pre>                   | Subscribe to notifications fired by our own queries |
| `notificationResult` | <pre>chan<- types.NotificationResult</pre> | A channel which receive notifications  |
| `onReconnectChannel` | <pre>chan<- interface{}</pre> | A channel which will be written to when reconnection will be triggered |

