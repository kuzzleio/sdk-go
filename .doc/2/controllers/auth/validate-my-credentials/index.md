---
code: true
type: page
title: ValidateMyCredentials
description: Validates the current user's credentials for the specified `<strategy>`.
---

# ValidateMyCredentials

Validates the current user's credentials for the specified `<strategy>`. The `result` field is `true` if the provided credentials are valid; otherwise an error is triggered. This route does not actually create or modify the user credentials. The credentials to send will depend on the authentication plugin and authentication strategy.

## Arguments

```go
func (a *Auth) ValidateMyCredentials(strategy string, credentials json.RawMessage, options types.QueryOptions) (bool, error)
```

| Arguments     | Type         | Description                                  |
| ------------- | ------------ | -------------------------------------------- |
| `strategy`    | <pre>string</pre>       | the strategy to use                          |
| `credentials` | <pre>string</pre>       | the new credentials                          |
| `options`     | <pre>QueryOptions</pre> | QueryOptions object containing query options |

### **Options**

Additional query options

| Property   | Type    | Description                       | Default |
| ---------- | ------- | --------------------------------- | ------- |
| `Queuable` | <pre>boolean</pre> | Make this request queuable or not | `true`  |

## Usage

<<< ./snippets/validate-my-credentials.go
