---
code: true
type: page
title: Send
description: Sends a query to the Kuzzle API
---

# Send

Sends a query to the [Kuzzle API](/core/1/api).

## Signature

```go
Send(query []byte, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse, requestId string) error
```

## Arguments

| Argument     | Type                                 | Description                 |
| ------------ | ------------------------------------ | --------------------------- |
| `query`      | <pre>[]byte</pre>        | API request                 |
| `options`    | <pre>QueryOptions</pre> | Additional query options    |
| `responseChannel` | chan<- \*types.KuzzleResponse | A channel to receive the API response | yes      |
| `requestId` | <pre>string</pre>        | Optional request identifier |

### **request**

Properties required for the Kuzzle API can be set in the [KuzzleRequest](https://github.com/kuzzleio/sdk-go/blob/master/types/kuzzle_request.go).
The following properties are the most common.

| Property     | Type         | Description                              | Required |
| ------------ | ------------ | ---------------------------------------- | -------- |
| `Controller` | string       | Controller name                          | yes      |
| `Action`     | string       | Action name                              | yes      |
| `Body`       | interface{}  | Query body for this action               | no       |
| `Index`      | string       | Index name for this action               | no       |
| `Collection` | string       | Collection name for this action          | no       |
| `Id`         | string       | id for this action                       | no       |
| `Volatile`   | VolatileData | Additional information to send to Kuzzle | no       |

### **options**

A [QueryOptions](https://github.com/kuzzleio/sdk-go/blob/master/types/query_options.go) containing additional query options
Theses properties can bet Get/Set.
The following properties are the most common.

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | bool | Make this request queuable or not | true    |

### **responseChannel**

A channel to receive the API response.
This channel will receive a [KuzzleResponse](https://github.com/kuzzleio/sdk-go/blob/master/types/kuzzle_response.go)

## Return

Return a [Kuzzle error](/sdk/go/1/essentials/error-handling) if the SDK can not connect to Kuzzle.