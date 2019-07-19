---
code: true
type: page
title: exists
description: Checks if collection exists
---

# Exists

Checks if a collection exists in Kuzzle.

## Arguments

```go
Exists(index string, collection string, options types.QueryOptions) (bool, error)
```

| Arguments    | Type               | Description
| ------------ | ------------------ | ---------------------------------- |
| `index`      | <pre>string</pre>             | Index name                         |
| `collection` | <pre>string</pre>             | Collection name                    |
| `options`    | <pre>types.QueryOptions</pre> | An object containing query options |

### **options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

True if the collection exists

## Usage

<<< ./snippets/exists.go
