---
code: true
type: page
title: mDelete
description: Deletes documents
---

# MDelete

Deletes multiple documents.

Throws a partial error (error code 206) if one or more document deletions fail.

The optional parameter `refresh` can be used with the value `wait_for` in order to wait for the document indexation (indexed documents are available for `search`).

## Arguments

```go
MDelete(
    index string,
    collection string,
    ids []string,
    options types.QueryOptions) ([]string, error)
```

<br/>

| Argument     | Type                          | Description                        |
| ------------ | ----------------------------- | ---------------------------------- |
| `index`      | <pre>string</pre>             | Index name                         |
| `collection` | <pre>string</pre>             | Collection name                    |
| `ids`        | <pre>[]string</pre>           | The ids of the documents to delete |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options  |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                        |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again       |
| `Refresh`  | <pre>string</pre><br/>(`""`)  | If set to `wait_for`, waits for the change to be reflected for `search` (up to 1s) |

## Return

Returns an array of strings containing the ids of the deleted documents.

## Usage

<<< ./snippets/m-delete.go
