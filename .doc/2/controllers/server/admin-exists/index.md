---
code: true
type: page
title: AdminExists
description: Returns information about Kuzzle server.
---

# AdminExists

Checks that an administrator account exists.

## Arguments

```go
func (s *Server) AdminExists(options types.QueryOptions) (bool, error)
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

Returns a `bool` set to `true` if an admin exists and `false` if it does not, or a `KuzzleError`. See how to [handle error](/sdk/go/1/essentials/error-handling).

## Usage

<<< ./snippets/admin-exists.go
