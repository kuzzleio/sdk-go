---
code: true
type: page
title: list
description: Lists the indexes
---

# List

Gets the complete list of indexes handled by Kuzzle.

## Arguments

```go
List(options types.QueryOptions) ([]string, error)
```

| Arguments | Type         | Description   |
| --------- | ------------ | ------------- |
| `options` | <pre>QueryOptions</pre> | Query options |

### **Options**

Additional query options

| Option     | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Returns an `Array` of strings containing the list of indexes names present in Kuzzle or an error

## Usage

<<< ./snippets/list.go
