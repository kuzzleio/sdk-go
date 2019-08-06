---
code: true
type: page
title: getMapping
description: Returns collection mapping
---

# GetMapping

Returns the mapping for the given `collection`.

## Arguments

```go
GetMapping(index string, collection string, options types.QueryOptions) (json.RawMessage, error)
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

Return a json representation of the mapping and an error is something was wrong.

## Usage

<<< ./snippets/get-mapping.go
