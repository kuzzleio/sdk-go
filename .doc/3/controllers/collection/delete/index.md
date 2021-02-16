---
code: true
type: page
title: delete
description: Deletes a collection
---

# delete

Deletes a collection.

<br/>

```go
Delete(index string, collection string, options types.QueryOptions) (json.RawMessage, error)
```

<br/>

| Arguments    | Type                    | Description     |
| ------------ | ----------------------- | --------------- |
| `index`      | <pre>string</pre>       | Index name      |
| `collection` | <pre>string</pre>       | Collection name |
| `options`    | <pre>QueryOptions</pre> | Query options   |


## Resolves

Resolves if the collection is successfully deleted.

## Usage

<<< ./snippets/delete.go
