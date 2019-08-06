---
code: true
type: page
title: list
description: Returns the collection list of an index
---

# List

Returns the complete list of realtime and stored collections in requested index sorted by name in alphanumerical order.
The `from` and `size` arguments allow pagination. They are returned in the response if provided.

## Arguments

```go
List(index string, options types.QueryOptions) (json.RawMessage, error)
```

| Arguments | Type               | Description                        |
| --------- | ------------------ | ---------------------------------- |
| `index`   | <pre>string</pre>             | Index name                         |
| `options` | <pre>types.QueryOptions</pre> | An object containing query options |

### **options**

Additional query options

| Property   | Type | Description                        | Default |
| ---------- | ---- | ---------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not  | `true`  |
| `from`     | <pre>int</pre>  | Offset of the first result         | `0`     |
| `size`     | <pre>int</pre>  | Maximum number of returned results | `10`    |

## Return

Return a json representation of the API return containing the collection list and an error is something was wrong.

## Usage

<<< ./snippets/list.go
