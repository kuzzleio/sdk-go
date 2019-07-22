---
code: true
type: page
title: getSpecifications
description: Returns the validation specifications
---

# GetSpecifications

Returns the validation specifications associated to the collection.

## Arguments

```go
GetSpecifications(index string, collection string, options types.QueryOptions) (json.RawMessage, error)
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

Return a json representation of the specifications and an error is something was wrong.

## Usage

<<< ./snippets/get-specifications.go
