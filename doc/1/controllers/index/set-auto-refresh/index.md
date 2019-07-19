---
code: true
type: page
title: setAutoRefresh
description: Sets the autorefresh flag
---

# setAutoRefresh(index, autorefresh, [options])

The setAutoRefresh action allows to set the autorefresh flag for the index.

Each index has an autorefresh flag.  
When set to true, each write request trigger a [refresh](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-refresh.html) action on Elasticsearch.  
Without a refresh after a write request, the documents may not be immediately visible in search.

:::info
A refresh operation comes with some performance costs.  
While forcing the autoRefresh can be convenient on a development or test environment,  
we recommend that you avoid using it in production or at least carefully monitor its implications before using it.
:::

## Arguments

```go
SetAutoRefresh(index string, autoRefresh bool, options types.QueryOptions) error
```

| Arguments     | Type         | Description      |
| ------------- | ------------ | ---------------- |
| `index`       | <pre>string</pre>       | Index name       |
| `autoRefresh` | <pre>Boolean</pre>      | autoRefresh flag |
| `options`     | <pre>QueryOptions</pre> | Query options    | no       |

### **Options**

Additional query options

| Option     | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Return an error or `nil`.

## Usage

<<< ./snippets/setAutoRefresh.go
