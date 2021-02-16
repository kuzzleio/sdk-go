---
code: true
type: page
title: Connect
description: Connect to kuzzle instance.
---

# Connect

Connects to kuzzle instance.

## Arguments

```go
Connect() (bool, error)
```

<br/>

| Argument   | Type     | Description      |
| ---------- | -------- | -------- |
| `event`    | <pre>int</pre> | Event constant from the `event` package |
| `channel` | <pre>channel</pre> | Event payload channel |

## Return

`true` and `nil ` if the connection is successful or `false` and an `error` otherwise.