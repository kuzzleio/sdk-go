---
code: true
type: page
title: mDelete
description: Deletes multiple indexes
---

# mDelete

Deletes multiple indexes at once.

## Arguments

```go
MDelete(indexes []string, options types.QueryOptions) ([]string, error)
```

| Arguments | Type         | Description                                   |
| --------- | ------------ | --------------------------------------------- |
| `indexes` | <pre>Array</pre>        | An array of strings containing indexes names. |
| `options` | <pre>QueryOptions</pre> | Query options                                 |

### **Options**

Additional query options

| Option     | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Returns an `Array` of strings containing the list of indexes names deleted or an error

## Usage

<<< ./snippets/mDelete.go
