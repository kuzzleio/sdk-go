---
code: true
type: page
title: RemoveListener
description: Removes a channel from an event
---

# RemoveListener

Removes a channel from an event.

## Arguments

```go
RemoveListener(event int, channel chan<- interface{})
```

<br/>

| Argument   | Type     | Description      |
| ---------- | -------- | -------- |
| `event`    | <pre>int</pre> | Event constant from the `event` package |
| `channel` | <pre>channel</pre> | Channel to unregister |

## Usage

<<< ./snippets/remove-listener.go
