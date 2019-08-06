---
code: true
type: page
title: deleteSpecifications
description: Deletes validation specifications for a collection
---

# DeleteSpecifications

Deletes the validation specifications associated with the collection.

## Arguments

```go
DeleteSpecifications(index string, collection string, options types.QueryOptions) error
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

Return an error or `nil` if collection successfully created.

## Usage

<<< ./snippets/delete-specifications.go
