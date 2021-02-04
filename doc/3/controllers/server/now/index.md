---
code: true
type: page
title: Now
description: Returns the current server timestamp, in Epoch-millis
---

# Now

Returns the current server timestamp, in Epoch-millis format.

## Arguments

```go
func (s *Server) Now(options types.QueryOptions) (int64, error)
```

| Arguments | Type               | Description    |
| --------- | ------------------ | -------------- |
| `options` | <pre>types.QueryOptions</pre> | Query options. |

### **Options**

Additional query options

| Option     | Type | Description                                                                  | Default |
| ---------- | ---- | ---------------------------------------------------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | If true, queues the request during downtime, until connected to Kuzzle again | `true`  |

## Return

Returns current server timestamp as `int64` or a `KuzzleError`. See how to [handle error](/sdk/go/1/essentials/error-handling).

## Usage

<<< ./snippets/now.go
