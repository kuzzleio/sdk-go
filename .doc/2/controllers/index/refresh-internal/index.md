---
code: true
type: page
title: refreshInternal
description: Forces refresh of Kuzzle internal index
---

# RefreshInternal

When writing or deleting security and internal documents (users, roles, profiles, configuration, etc.) in Kuzzle, the update needs to be indexed before being reflected in the search index.

The `refreshInternal` action forces a [refresh](/sdk/go/1/controllers/index/refresh/), on the internal index, making the documents available to search immediately.

::: info
A refresh operation comes with some performance costs.

From [Elasticsearch documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-refresh.html):
"While a refresh is much lighter than a commit, it still has a performance cost. A manual refresh can be useful when writing tests, but don’t do a manual refresh every time you index a document in production; it will hurt your performance. Instead, your application needs to be aware of the near real-time nature of Elasticsearch and make allowances for it."
:::

## Arguments

```go
RefreshInternal(options types.QueryOptions) error
```

<br/>

| Arguments | Type         | Description   |
| --------- | ------------ | ------------- |
| `options` | <pre>QueryOptions</pre> | Query options |

### Options

The `options` arguments can contain the following option properties:

| Option     | Type (default) | Description                       |
| ---------- | -------------- | --------------------------------- |
| `queuable` | <pre>bool (true)</pre> | If true, queues the request during downtime, until connected to Kuzzle again |

## Return

Return an error or `nil` if index successfully refreshed.

## Usage

<<< ./snippets/refreshInternal.go
