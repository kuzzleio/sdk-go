---
code: true
type: page
title: truncate
description: Removes all documents from collection
---

# Truncate

Removes all documents from a collection while keeping the associated mapping.  
It is faster than deleting all documents from a collection.

## Arguments

```go
Truncate(index string, collection string, options types.QueryOptions) error
```

| Arguments    | Type               | Description     |
| ------------ | ------------------ | --------------- |
| `index`      | <pre>string</pre>             | Index name      |
| `collection` | <pre>string</pre>             | Collection name |
| `options`    | <pre>types.QueryOptions</pre> | Query options   |

### **options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Return an error is something was wrong.

## Usage

<<< ./snippets/truncate.go
