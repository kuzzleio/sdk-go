---
code: true
type: page
title: create
description: Creates an index
---

# Create

Create a new index in Kuzzle

## Arguments

```go
Create(index string, options types.QueryOptions) error
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

Return an error or `nil` if index successfully created.

## Usage

<<< ./snippets/create.go
