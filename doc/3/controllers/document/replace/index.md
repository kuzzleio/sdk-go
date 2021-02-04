---
code: true
type: page
title: replace
description: Replaces a document
---

# Replace

Replaces the content of an existing document.

## Arguments

```go
Replace(
  index string,
  collection string,
  id string,
  document json.RawMessage,
  options types.QueryOptions) (json.RawMessage, error)
```

<br/>

| Argument     | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `id`         | <pre>string</pre>             | Document ID                       |
| `document`   | <pre>string</pre>             | Document body                     |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                        |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again       |
| `Refresh`  | <pre>string</pre><br/>(`""`)  | If set to `wait_for`, waits for the change to be reflected for `search` (up to 1s) |

## Return

Returns a json.RawMessage containing the replaced document.

## Usage

<<< ./snippets/replace.go
