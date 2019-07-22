---
code: true
type: page
title: playQueue
description: Plays the offline request queue
---

# PlayQueue

Plays the requests queued during `offline` state.  
Works only if the SDK is not in a `offline` state, and if the `autoReplay` option is set to `false`.

## Arguments

```go
PlayQueue()
```

## Usage

<<< ./snippets/play-queue.go
