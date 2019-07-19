---
code: true
type: page
title: RemoveAllListeners
description: Removes all channels, or all channels from an event
---

# RemoveAllListeners

Removes all channels from an event.  
If no eventName is specified, removes all channels from all events.

Channels removed this way are **not** closed.

## Arguments

```go
RemoveAllListeners(event int)
```

<br/>

| Argument   | Type     | Description      |
| ---------- | -------- | -------- |
| `event`    | <pre>int</pre> | Event constant from the `event` package |

## Usage

<<< ./snippets/remove-all-listeners.go
