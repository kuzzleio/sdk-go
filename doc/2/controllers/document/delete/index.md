---
code: true
type: page
title: delete
description: Deletes a document from kuzzle
---

# Delete

Deletes a document.

The optional parameter `refresh` can be used with the value `wait_for` in order to wait for the document to be indexed (and to no longer be available in search).

## Arguments

```go
Delete(
  index string,
  collection string,
  _id string,
  options types.QueryOptions) (string, error)
```

<br/>

| Argument     | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `id`         | <pre>string</pre>             | Document ID                       |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                        |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again       |
| `Refresh`  | <pre>string</pre><br/>(`""`)  | If set to `wait_for`, waits for the change to be reflected for `search` (up to 1s) |

## Return

Returns the id of the deleted document.

## Usage

<<< ./snippets/delete.go
