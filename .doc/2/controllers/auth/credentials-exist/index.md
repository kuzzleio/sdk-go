---
code: true
type: page
title: CredentialsExist
description: Checks that the current user has credentials for the specified strategy
---

# CredentialsExist

Checks that the current user has credentials for the specified strategy.

## Arguments

```go
func (a *Auth) CredentialsExist(strategy string, options types.QueryOptions) (bool, error)
```

| Arguments  | Type         | Description                          | Required |
| ---------- | ------------ | ------------------------------------ | -------- |
| `strategy` | <pre>string</pre>       | Strategy to use                      | yes      |
| `options`  | <pre>QueryOptions</pre> | A structure containing query options | yes      |

### **Options**

Additional query options

| Property   | Type    | Description                       | Default |
| ---------- | ------- | --------------------------------- | ------- |
| `Queuable` | <pre>boolean</pre> | Make this request queuable or not | `true`  |

## Return

True if exists, false if not.

## Usage

<<< ./snippets/credentials-exist.go
