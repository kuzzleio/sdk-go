---
code: true
type: page
title: refresh
description: Forces Elasticsearch search index update
---

# Refresh

When writing or deleting documents in Kuzzle, the update needs to be indexed before being available in search results.

:::info
A refresh operation comes with some performance costs.

From [Elasticsearch documentation](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/docs-refresh.html):
> "While a refresh is much lighter than a commit, it still has a performance cost. A manual refresh can be useful when writing tests, but don’t do a manual refresh every time you index a document in production; it will hurt your performance. Instead, your application needs to be aware of the near real-time nature of Elasticsearch and make allowances for it."
:::

## Arguments

```go
Refresh(index string, options types.QueryOptions) error
```

| Arguments | Type         | Description   |
| --------- | ------------ | ------------- |
| `index`   | <pre>string</pre>       | Index name    |
| `options` | <pre>QueryOptions</pre> | Query options |

### **Options**

Additional query options

| Option     | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Return an error or `nil` if index successfully refreshed.

## Usage

<<< ./snippets/refresh.go
