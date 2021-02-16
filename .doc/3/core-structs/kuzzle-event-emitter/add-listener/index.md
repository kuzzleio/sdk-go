---
code: true
type: page
title: AddListener
description: Adds a new channel for an event
---

# AddListener

Adds a channel at the end of list of registered channels for that event. 
Whenever an event is triggered, registered channels are fed in the order they were registered.

## Arguments

```js
AddListener(event int, channel chan<- interface{})
```

<br/>

| Argument   | Type     | Description      |
| ---------- | -------- | -------- |
| `event`    | <pre>int</pre> | Event constant from the `event` package |
| `channel` | <pre>channel</pre> | Event payload channel |

## Usage

<<< ./snippets/add-listener.go
