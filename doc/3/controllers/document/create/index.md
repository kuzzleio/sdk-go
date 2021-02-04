---
code: true
type: page
title: create
description: Creates a new document
---

# Create

Creates a new document in the persistent data storage.

Returns an error if the document already exists.

The optional parameter `refresh` can be used with the value `wait_for` in order to wait for the document to be indexed (indexed documents are available for `search`).

## Arguments

```go
Create(
  index string,
  collection string,
  id string,
  document json.RawMessage,
  options types.QueryOptions) (json.RawMessage, error) {
```

<br/>

| Argument     | Type                          | Description                                                                  |
| ------------ | ----------------------------- | ---------------------------------------------------------------------------- |
| `index`      | <pre>string</pre>             | Index name                                                                   |
| `collection` | <pre>string</pre>             | Collection name                                                              |
| `id`         | <pre>string</pre>             | Optional document id. If set to a blank string, will use a auto-generated id |
| `document`   | <pre>json.RawMessage</pre>    | Document content                                                             |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options                                            |

### options

Additional query options

| Option     | Type<br/>(default)           | Description                                                                        |
| ---------- | ---------------------------- | ---------------------------------------------------------------------------------- |
| `Queuable` | <pre>bool</pre><br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again       |
| `Refresh`  | <pre>string</pre><br/>(`""`) | If set to `wait_for`, waits for the change to be reflected for `search` (up to 1s) | `` |

## Return

Returns a json.RawMessage containing the document creation result.

| Name      | Type              | Description                                                                      |
| --------- | ----------------- | -------------------------------------------------------------------------------- |
| \_id      | string            | Newly created document ID                                                        |
| \_version | int               | Version of the document in the persistent data storage                           |
| \_source  | object            | The created document                                                             |
| result    | <pre>string</pre> | Set to `created` in case of success and `updated` if the document already exists |

## Usage

<<< ./snippets/create.go
