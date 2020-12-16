---
code: true
type: page
title: deleteByQuery
description: Delete documents matching query
---

# DeleteByQuery

Deletes documents matching the provided search query.

Kuzzle uses the [ElasticSearch Query DSL](https://www.elastic.co/guide/en/elasticsearch/reference/7.4/query-dsl.html) syntax.

<SinceBadge version="change-me"/>

This method also supports the [Koncorde Filters DSL](/core/2/guides/cookbooks/realtime-api) to match documents by passing the `lang` argument with the value `koncorde`.  
Koncorde filters will be translated into an Elasticsearch query.  

::: warning
Koncorde `bool` operator and `regexp` clause are not supported for search queries.
:::

## Arguments

```go
DeleteByQuery(
  index string,
  collection string,
  query json.RawMessage,
  options types.QueryOptions) ([]string, error)
```

<br/>

| Argument     | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `query`      | <pre>string</pre>             | Query to match                    |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                        |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again       |
| `Refresh`  | <pre>string</pre><br/>(`""`)  | If set to `wait_for`, waits for the change to be reflected for `search` (up to 1s) |
| `Lang`     | <pre>string</pre>               | Specify the query language to use. By default, it's `elasticsearch` but `koncorde` can also be used. <SinceBadge version="change-me"/> |

## Return

Returns the list of the deleted document ids.

## Usage

<<< ./snippets/delete-by-query.go
