---
code: true
type: page
title: Publish
description: Publishes a real-time message
---

# Publish

Sends a real-time `<message>` to Kuzzle. The `<message>` will be dispatched to all clients with subscriptions matching the `<index>`, the `<collection>` and the `<message>` content.

The `<index>` and `<collection>` are indicative and serve only to distinguish the rooms. They are not required to exist in the database

**Note:** real-time messages are not persisted in the database.

## Arguments

```go
func (r *Realtime) Publish(
  index string,
  collection string,
  message json.RawMessage,
  options types.QueryOptions
) error
```

<br/>

| Arguments    | Type                          | Description     |
| ------------ | ----------------------------- | --------------- |
| `index`      | <pre>string</pre>             | Index name      |
| `collection` | <pre>string</pre>             | Collection name |
| `message`    | <pre>json.RawMessage</pre>    | Message to send |
| `options`    | <pre>types.QueryOptions</pre> | Query options   |

### options

Additional query options

| Option     | Type<br/>(default)           | Description                       |
| ---------- | ---------------------------- | --------------------------------- |
| `queuable` | <pre>bool</pre><br/>(`true`) | Make this request queuable or not |

## Return

Return an error is something was wrong.

## Usage

<<< ./snippets/publish.go
