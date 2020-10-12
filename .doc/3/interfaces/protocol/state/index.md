---
code: true
type: page
title: State
description: Gets the current state
---

# State

Gets the current connection state.

## Signature

```cpp
State() int
```

## Return

The current connection state, values can be from the State enum:

```go
Connecting
Disconnected
Connected
Initializing
Ready
Logged_out
Error
Offline
```
