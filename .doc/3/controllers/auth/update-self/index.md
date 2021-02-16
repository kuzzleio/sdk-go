---
code: true
type: page
title: UpdateSelf
description: Updates the current user object in Kuzzle.
---

# UpdateSelf

Updates the current user object in Kuzzle.

## Arguments

```go
func (a *Auth) UpdateSelf(data json.RawMessage, options types.QueryOptions) (*security.User, error)
```

| Arguments | Type         | Description                                  |
| --------- | ------------ | -------------------------------------------- |
| `content` | <pre>string</pre>       | the new credentials                          |
| `options` | <pre>QueryOptions</pre> | QueryOptions object containing query options |

### **Options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

A pointer to a security.User object and an error or `nil`

## Usage

<<< ./snippets/update-self.go
