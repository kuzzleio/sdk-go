---
code: true
type: page
title: exists
description: Checks for index existence
---

# Exists

Checks if the given index exists in Kuzzle.

## Arguments

```go
Exists(index string, options types.QueryOptions) (bool, error)
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

Returns a `bool` that indicate whether the index exists, or an error

## Usage

<<< ./snippets/exists.go
