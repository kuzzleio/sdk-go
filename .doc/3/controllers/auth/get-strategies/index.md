---
code: true
type: page
title: GetStrategies
description: Gets all authentication strategies registered in Kuzzle.
---

# GetStrategies

Gets all authentication strategies registered in Kuzzle.

## Arguments

```go
func (a *Auth) GetStrategies(options types.QueryOptions) ([]string, error)
```

| Arguments | Type            | Description                                                       |
| --------- | --------------- | ----------------------------------------------------------------- |
| `options` | <pre>query_options\*</pre> | A pointer to a `kuzzleio::query_options` containing query options |

### **Options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

An array of string containing the list of strategies and an error or nil.

## Usage

<<< ./snippets/get-strategies.go
