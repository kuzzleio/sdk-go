---
code: true
type: page
title: GetLastStats
description: Returns the most recent statistics snapshot.
---

# GetLastStats

Returns the most recent statistics snapshot.
By default, snapshots are made every 10 seconds and they are stored for 1 hour.

These statistics include:

- the number of connected users per protocol (not available for all protocols)
- the number of ongoing requests
- the number of completed requests since the last frame
- the number of failed requests since the last frame

## Arguments

```go
func (s *Server) GetLastStats(options types.QueryOptions) (json.RawMessage, error)
```

| Arguments | Type               | Description                         |
| --------- | ------------------ | ----------------------------------- |
| `options` | <pre>types.QueryOptions</pre> | An object containing query options. |

### **Options**

Additional query options

| Option     | Type | Description                                                                  | Default |
| ---------- | ---- | ---------------------------------------------------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | If true, queues the request during downtime, until connected to Kuzzle again | `true`  |

## Return

Returns the most recent statistics snapshot as a `json.RawMessage` or a `KuzzleError`. See how to [handle error](/sdk/go/1/essentials/error-handling).

## Return

## Usage

<<< ./snippets/get-last-stats.go
