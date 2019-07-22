---
code: true
type: page
title: UpdateMyCredentials
description: Updates the current user's credentials for the specified strategy.
---

# UpdateMyCredentials

Updates the current user's credentials for the specified `<strategy>`. The credentials to send will depend on the authentication plugin and the authentication strategy.

## Arguments

```go
func (a *Auth) UpdateMyCredentials(strategy string, credentials json.RawMessage, options types.QueryOptions) (json.RawMessage, error)
```

| Arguments     | Type            | Description                                  |
| ------------- | --------------- | -------------------------------------------- |
| `strategy`    | <pre>string</pre>          | the strategy to use                          |
| `credentials` | <pre>json.RawMessage</pre> | the new credentials                          |
| `options`     | <pre>QueryOptions</pre>    | QueryOptions object containing query options |

### **Options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

A JSON representing the new credentials and an error or `nil`.

## Usage

<<< ./snippets/update-my-credentials.go
