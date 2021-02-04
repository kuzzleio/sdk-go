---
code: true
type: page
title: delete
description: Deletes an index
---

# Delete

Deletes an entire index from Kuzzle.

## Arguments

```go
Delete(index string, options types.QueryOptions) error
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

Returns an error or `nil` if the request succeed.

## Usage

<<< ./snippets/delete.go
