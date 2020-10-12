---
code: true
type: page
title: Logout
description: Revokes the user's token & unsubscribe them from registered rooms.
---

# Logout

Revokes the user's token & unsubscribe them from registered rooms.

## Arguments

```go
func (a *Auth) Logout() error
```

## Return

Return an error or `nil` if the credentials are successfully deleted

## Usage

<<< ./snippets/logout.go
