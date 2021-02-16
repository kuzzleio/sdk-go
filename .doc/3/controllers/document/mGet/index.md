---
code: true
type: page
title: mGet
description: Gets multiple documents
---

# MGet

Gets multiple documents.

Returns a partial error (error code 206) if one or more document can not be retrieved.

## Arguments

```go
MGet(
    index string,
    collection string,
    ids []string,
    options types.QueryOptions) (json.RawMessage, error)
```

<br/>

| Arguments    | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `ids`        | <pre>[]string</pre>           | Document IDs                      |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                  |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again |

## Return

Returns a json.RawMessage containing two arrays, successes and errors.

The `successes` array contain the list of retrieved documents.

Each document have the the following properties:
| Name       | Type                       | Description                                            |
| ---------- | -------------------------- | ------------------------------------------------------ |
| `_id`      | <pre>string</pre>          | Document ID                                            |
| `_version` | <pre>int</pre>             | Version of the document in the persistent data storage |
| `_source`  | <pre>json.RawMessage</pre> | Document content                                       |

The `errors` array contain the IDs of not found documents.

## Usage

<<< ./snippets/m-get.go
