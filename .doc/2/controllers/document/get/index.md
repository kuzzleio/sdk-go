---
code: true
type: page
title: get
description: Get a document from kuzzle
---

# Get

Gets a document.

## Arguments

```go
Get(
  index string,
  collection string,
  _id string,
  options types.QueryOptions) (json.RawMessage, error)
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

| Option     | Type<br/>(default)            | Description                                                                  |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again |

## Return

Returns a json.RawMessage containing the document.

| Name      | Type   | Description                                            |
| --------- | ------ | ------------------------------------------------------ |
| \_id      | string | Newly created document ID                              |
| \_version | int    | Version of the document in the persistent data storage |
| \_source  | object | The created document                                   |

## Usage

<<< ./snippets/get.go
