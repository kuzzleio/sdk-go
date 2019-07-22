---
code: true
type: page
title: Subscribe
description: Subscribes to real-time notifications
---

# Subscribe

Subscribes by providing a set of filters: messages, document changes and, optionally, user events matching the provided filters will generate [real-time notifications](/core/1/api/essentials/notifications/), sent to you in real-time by Kuzzle.

## Arguments

```go
func (r *Realtime) Subscribe(
  index string,
  collection string,
  filters json.RawMessage,
  listener chan<- types.NotificationResult,
  options types.RoomOptions
) (*types.SubscribeResult, error)
```

<br/>

| Arguments    | Type                                       | Description                                                     |
| ------------ | ------------------------------------------ | --------------------------------------------------------------- |
| `index`      | <pre>string</pre>                          | Index name                                                      |
| `collection` | <pre>string</pre>                          | Collection name                                                 |
| `filters`    | <pre>json.RawMessage</pre>                 | A set of filters following [Koncorde syntax](/core/1/guides/cookbooks/realtime-api/) |
| `listener`   | <pre>chan<- types.NotificationResult</pre> | Channel receiving the notification                              |
| `options`    | <pre>types.RoomOptions</pre>               | A struct containing subscription options                        |

### listener

A channel for [types.NotificationResult](/sdk/go/1/essentials/realtime-notifications) objects.
The channel will receive an object each time a new notifications is received.

### options

Additional subscription options.

| Property          | Type<br/>(default)                    | Description                                                                                              |
| ----------------- | ------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `scope`           | <pre>string</pre><br/>(`all`)         | Subscribe to document entering or leaving the scope</br>Possible values: `all`, `in`, `out`, `none`      |
| `users`           | <pre>string</pre><br/>(`none`)        | Subscribe to users entering or leaving the room</br>Possible values: `all`, `in`, `out`, `none`          |
| `subscribeToSelf` | <pre>bool</pre><br/>(`true`)          | Subscribe to notifications fired by our own queries                                                      |
| `volatile`        | <pre>json.RawMessage</pre><br/>(`{}`) | subscription information, used in [user join/leave notifications](/core/1/api/essentials/volatile-data/) |

## Return

Return an error if something was wrong or a `types.SubscribeResult` containing the following properties:

| Property  | Type              | Description    |
| --------- | ----------------- | -------------- |
| `Room`    | <pre>string</pre> | The room ID    |
| `Channel` | <pre>string</pre> | The channel ID |

## Usage

_Simple subscription to document notifications_

<<< ./snippets/document-notifications.go

_Subscription to document notifications with scope option_

<<< ./snippets/document-notifications-leave-scope.go

_Subscription to message notifications_

<<< ./snippets/message-notifications.go

_Subscription to user notifications_

<<< ./snippets/user-notifications.go
