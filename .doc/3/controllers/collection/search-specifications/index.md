---
code: true
type: page
title: searchSpecifications
description: Searches collection specifications
---

# SearchSpecifications

Searches collection specifications.

There is a limit to how many items can be returned by a single search query.
That limit is by default set at 10000, and you can't get over it even with the from and size pagination options.

:::info
When processing a large number of items (i.e. more than 1000), it is advised to paginate the results using [SearchResult.next](/sdk/go/2/core-structs/search-result#methods) rather than increasing the size parameter.
:::

## Arguments

```go
SearchSpecifications(
  query json.RawMessage,
  options types.QueryOptions) (*types.SearchResult, error)
```

<br/>

| Argument  | Type                          | Description                       |
| --------- | ----------------------------- | --------------------------------- |
| `query`   | <pre>json.RawMessage</pre>    | Query to match                    |
| `options` | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

| Options    | Type (default)               | Description                                                                                                                                                                                                       |
| ---------- | ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `queuable` | <pre>boolean</pre> (`true`)  | If true, queues the request during downtime, until connected to Kuzzle again                                                                                                                                      |
| `from`     | <pre>int</pre><br/>(`0`)     | Offset of the first document to fetch                                                                                                                                                                             |
| `size`     | <pre>int</pre><br/>(`10`)    | Maximum number of documents to retrieve per page                                                                                                                                                                  |
| `scroll`   | <pre>string</pre><br/>(`""`) | When set, gets a forward-only cursor having its ttl set to the given value (ie `30s`; cf [elasticsearch time limits](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/common-options.html#time-units)) |

## Query properties

### Optional:

- `query`: the search query itself, using the [ElasticSearch Query DSL](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/query-dsl.html) syntax.
- `aggregations`: controls how the search results should be [aggregated](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/search-aggregations.html)
- `sort`: contains a list of fields, used to [sort search results](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/search-request-sort.html), in order of importance.

An empty body matches all documents in the queried collection.

## Return

Returns a [types.SearchResult](/sdk/go/1/core-structs/search-result) struct

## Usage

<<< ./snippets/search-specifications.go
