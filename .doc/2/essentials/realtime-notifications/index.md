---
code: false
type: page
title: Realtime notifications
description: List of realtime notifications sent by Kuzzle
order: 100
---

# Notifications

The [Realtime.Subscribe](/sdk/go/1/controllers/realtime/subscribe/) method takes a channel for `types.NotificationResult` objects, whose content depend on the type of notification received.

## Document & messages

These notifications represent [documents changes & messages](/core/1/api/essentials/notifications#documents-changes-messages).

| Property     | Type                       | Description                                                                                           |
| ------------ | -------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Action`     | string                     | API controller's action                                                                               |
| `Collection` | string                     | Data collection                                                                                       |
| `Controller` | string                     | API controller                                                                                        |
| `Index`      | string                     | Data index                                                                                            |
| `Protocol`   | string                     | Network protocol used to modify the document                                                          |
| `Result`     | \*types.NotificationResult | Notification content                                                                                  |
| `RoomId`     | string                     | Subscription channel identifier. Can be used to link a notification to its corresponding subscription |
| `Scope`      | string                     | `in`: document enters (or stays) in the scope<br/>`out`: document leaves the scope                    |
| `Timestamp`  | int                        | Timestamp of the event, in Epoch-millis format                                                        |
| `Type`       | string                     | `document`: the notification type                                                                     |
| `Volatile`   | json.RawMessage            | Request [volatile data](/core/1/api/essentials/volatile-data/)                                        |

The `Result` property has the following structure for document notifications & messages:

| Property  | Type            | Description                                                                                           |
| --------- | --------------- | ----------------------------------------------------------------------------------------------------- |
| `id`      | string          | Document unique ID<br/>`null` if the notification is from a real-time message                         |
| `content` | json.RawMessage | A JSON String message or full document content. Not present if the event is about a document deletion |

## User

These notifications represent [user events](/core/1/api/essentials/notifications#user-notification).

| Property     | Type                       | Description                                                                                           |
| ------------ | -------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Action`     | string                     | API controller's action                                                                               |
| `Collection` | string                     | Data collection                                                                                       |
| `Controller` | string                     | API controller                                                                                        |
| `Index`      | string                     | Data index                                                                                            |
| `Protocol`   | string                     | Network protocol used by the entering/leaving user                                                    |
| `Result`     | \*types.NotificationResult | Notification content                                                                                  |
| `RoomId`     | string                     | Subscription channel identifier. Can be used to link a notification to its corresponding subscription |
| `Timestamp`  | int                        | Timestamp of the event, in Epoch-millis format                                                        |
| `Type`       | string                     | `user`: the notification type                                                                         |
| `User`       | string                     | `in`: a new user has subscribed to the same filters<br/>`out`: a user cancelled a shared subscription |
| `Volatile`   | json.RawMessage            | Request [volatile data](/core/1/api/essentials/volatile-data/)                                        |

The `Result` property has the following structure for user events:

| Property | Type | Description                                        |
| -------- | ---- | -------------------------------------------------- |
| `count`  | int  | Updated users count sharing that same subscription |
