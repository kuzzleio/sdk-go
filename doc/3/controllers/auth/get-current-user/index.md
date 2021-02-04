---
code: true
type: page
title: GetCurrentUser
description: Returns the profile object for the user linked to the `JSON Web Token`
---

# GetCurrentUser

Returns the profile object for the user linked to the `JSON Web Token`, provided in the query or the `Authorization` header.

## Arguments

```go
func (a *Auth) GetCurrentUser() (*security.User, error)
```

## Return

A pointer to security.User object containing:

| Property     | Type                   | Description                         |
| ------------ | ---------------------- | ----------------------------------- |
| `Id`         | <pre>string</pre>                 | The user ID                         |
| `Content`    | <pre>map[string]interface{}</pre> | The user content                    |
| `ProfileIds` | <pre>[]string</pre>    | An array containing the profile ids |

## Usage

<<< ./snippets/get-current-user.go
