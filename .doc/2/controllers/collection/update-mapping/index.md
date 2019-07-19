---
code: true
type: page
title: updateMapping
description: Update the collection mapping
---

# UpdateMapping

Update the collection mapping.
Mapping allow you to exploit the full capabilities of our
persistent data storage layer, [ElasticSearch](https://www.elastic.co/products/elasticsearch) (check here the [mapping capabilities of ElasticSearch](https://www.elastic.co/guide/en/elasticsearch/reference/5.4/mapping.html)).

## Arguments

```go
UpdateMapping(index string, collection string, mapping json.RawMessage, options types.QueryOptions) error
```

| Arguments    | Type            | Description                            |
| ------------ | --------------- | -------------------------------------- |
| `index`      | <pre>string</pre>          | Index name                             |
| `collection` | <pre>string</pre>          | Collection name                        |
| `mapping`    | <pre>json.RawMessage</pre> | Collection data mapping in JSON format |
| `options`    | <pre>QueryOptions</pre>    | Query options                          |

### **mapping**

An string containing the JSON representation of the collection data mapping.

The mapping must have a root field `properties` that contain the mapping definition:

```json
{
  "properties": {
    "field1": { "type": "text" },
    "field2": {
      "properties": {
        "nestedField": { "type": "keyword" }
      }
    }
  }
}
```

More informations about database mappings [here](/core/1/guides/essentials/database-mappings).

### **options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Return an error if something went wrong.

## Usage

<<< ./snippets/update-mapping.go
