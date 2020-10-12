---
code: true
type: page
title: validateSpecifications
description: Validates specifications format
---

# ValidateSpecifications

The validateSpecifications method checks if a validation specification is well formatted. It does not store nor modify the existing specification.

When the validation specification is not formatted correctly, a detailed error message is returned to help you to debug.

## Arguments

```go
ValidateSpecifications(index string, collection string, specifications json.RawMessage, options types.QueryOptions) (types.ValidationResponse, error)
```

| Arguments        | Type            | Description                            |
| ---------------- | --------------- | -------------------------------------- |
| `index`          | <pre>string</pre>          | Index name                             |
| `collection`     | <pre>string</pre>          | Collection name                        |
| `specifications` | <pre>json.RawMessage</pre> | Collection data mapping in JSON format |
| `options`        | <pre>QueryOptions</pre>    | Query options                          |

### **specifications**

A JSON representation of the specifications.

The JSON must follow the [Specification Structure](/core/1/guides/cookbooks/datavalidation):

```json
{
  "myindex": {
    "mycollection": {
      "strict": "<boolean>",
      "fields": {
        // ... specification for each field
      }
    }
  }
}
```

### **options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

A `types.ValidationResponse` which contain information about the specifications validity.

| Property      | Type                | Description                             |
| ------------- | ------------------- | --------------------------------------- |
| `Valid`       | bool                | Specification validity                  |
| `Details`     | <pre>[]string</pre> | Details about each specification errors |
| `Description` | string              | General error message                   |

## Usage

<<< ./snippets/validate-specifications.go
