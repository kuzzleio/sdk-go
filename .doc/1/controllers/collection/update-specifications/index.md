---
code: true
type: page
title: updateSpecifications
description: Updates the validation specifications
---

# UpdateSpecifications

The updateSpecifications method allows you to create or update the validation specifications for one or more index/collection pairs.

When the validation specification is not formatted correctly, a detailed error message is returned to help you to debug.

## Arguments

```go
UpdateSpecifications(index string, collection string, specifications json.RawMessage, options types.QueryOptions) (json.RawMessage, error)
```

| Arguments        | Type            | Description                   |
| ---------------- | --------------- | ----------------------------- |
| `index`          | <pre>string</pre>          | Index name                    |
| `collection`     | <pre>string</pre>          | Collection name               |
| `specifications` | <pre>json.RawMessage</pre> | Specifications in JSON format |
| `options`        | <pre>QueryOptions</pre>    | Query options                 |

### **specifications**

A JSON representation of the specifications.

The JSON must follow the [Specification Structure](/core/1/guides/cookbooks/datavalidation):

```json
{
  "strict": "<boolean>",
  "fields": {
    // ... specification for each field
  }
}
```

### **options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Return a JSON representation of the specifications.
Return an error with a global description of errors.

## Usage

<<< ./snippets/update-specifications.go
