---
code: true
type: page
title: DeleteMyCredentials
description: Deletes the current user's credentials for the specified strategy
---

# DeleteMyCredentials

Deletes the current user's credentials for the specified `<strategy>`. If the credentials that generated the current JWT are removed, the user will remain logged in until he logs out or his session expires, after that they will no longer be able to log in with the deleted credentials.

## Arguments

```go
func (a *Auth) DeleteMyCredentials(strategy string, options types.QueryOptions) error
```

| Arguments  | Type         | Description                                  | Required |
| ---------- | ------------ | -------------------------------------------- | -------- |
| `strategy` | <pre>string</pre>       | the strategy to use                          | yes      |
| `options`  | <pre>QueryOptions</pre> | QueryOptions object containing query options | yes      |

### **Options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

Return an error or `nil` if the credentials are successfully deleted

## Usage

<<< ./snippets/delete-my-credentials.go
