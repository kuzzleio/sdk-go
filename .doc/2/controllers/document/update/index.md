---
code: true
type: page
title: update
description: Updates a document
---

# Update

Updates a document content.

Conflicts may occur if the same document gets updated multiple times within a short timespan, in a database cluster.
You can set the `retryOnConflict` optional argument (with a retry count), to tell Kuzzle to retry the failing updates the specified amount of times before rejecting the request with an error.

## Arguments

```go
Update(
  index string,
  collection string,
  id string,
  document json.RawMessage,
  options types.QueryOptions) (json.RawMessage, error)
```

<br/>

| Argument     | Type                          | Description                       |
| ------------ | ----------------------------- | --------------------------------- |
| `index`      | <pre>string</pre>             | Index name                        |
| `collection` | <pre>string</pre>             | Collection name                   |
| `id`         | <pre>string</pre>             | Document ID                       |
| `document`   | <pre>string</pre>             | Document body                     |
| `options`    | <pre>types.QueryOptions</pre> | A struct containing query options |

### options

Additional query options

| Option            | Type<br/>(default)            | Description                                                                        |
| ----------------- | ----------------------------- | ---------------------------------------------------------------------------------- |
| `Queuable`        | <pre>bool</pre> <br/>(`true`) | If true, queues the request during downtime, until connected to Kuzzle again       |
| `Refresh`         | <pre>string</pre><br/>(`""`)  | If set to `wait_for`, waits for the change to be reflected for `search` (up to 1s) |
| `RetryOnConflict` | <pre>int</pre><br/>(`0`)      | Number of times the database layer should retry in case of version conflict        |

## Return

Returns a json.RawMessage containing the document update result.

| Name      | Type              | Description                                            |
| --------- | ----------------- | ------------------------------------------------------ |
| \_id      | <pre>string</pre> | Newly created document ID                              |
| \_version | int               | Version of the document in the persistent data storage |
| result    | <pre>string</pre> | Set to `updated` in case of success                    |

## Usage

<<< ./snippets/update.go
