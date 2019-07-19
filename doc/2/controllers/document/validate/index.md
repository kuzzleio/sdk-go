---
code: true
type: page
title: validate
description: Validates a document
---

# Validate

Validates data against existing validation rules.

Note that if no validation specifications are set for the `<index>`/`<collection>`, the document will always be valid.

This request does **not** store or publish the document.

## Arguments

```go
Validate(
  index string,
  collection string,
  document json.RawMessage,
  options types.QueryOptions) (bool, error)
```

<br/>

| Argument     | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `document`   | <pre>string</pre>             | Document body                     |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option     | Type<br/>(default)            | Description                                                                  |
| ---------- | ----------------------------- | ---------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again |

## Return

Returns a boolean value set to true if the document is valid and false otherwise.

## Usage

<<< ./snippets/validate.go
