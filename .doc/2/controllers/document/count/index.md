---
code: true
type: page
title: count
description: Counts documents matching the given query
---

# Count

Counts documents in a collection.

A query can be provided to alter the count result, otherwise returns the total number of documents in the collection.

Kuzzle uses the [ElasticSearch Query DSL](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/query-dsl.html) syntax.

## Arguments

```go
Count(
  index string,
  collection string,
  query json.RawMessage,
  options types.QueryOptions) (int, error)
```

<br/>

| Argument     | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `query`      | <pre>json.RawMessage</pre>    | Query to match                    |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                  |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again |

## Return

Returns the number of documents matching the given query.

## Usage

<<< ./snippets/count.go
