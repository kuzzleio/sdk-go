---
code: true
type: page
title: query
description: Base method to send API query to Kuzzle
---

# Query

Base method used to send queries to Kuzzle, following the [API Documentation](/core/1/api).

:::warning
This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
:::

## Arguments

```go
Query(request *types.KuzzleRequest, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse)
```

| Argument          | Type                          | Description                           |
| ----------------- | ----------------------------- | ------------------------------------- |
| `request`         | <pre>\*types.KuzzleRequest</pre>         | API request options                   |
| `options`         | <pre>types.QueryOptions</pre>            | Additional query options              |
| `responseChannel` | <pre>chan<- \*types.KuzzleResponse</pre> | A channel to receive the API response |

### **request**

Properties required for the Kuzzle API can be set in the [KuzzleRequest](https://github.com/kuzzleio/sdk-go/blob/master/types/kuzzle_request.go).
The following properties are the most common.

| Property     | Type         | Description                              |
| ------------ | ------------ | ---------------------------------------- |
| `Controller` | <pre>string</pre>       | Controller name                          |
| `Action`     | <pre>string</pre>       | Action name                              |
| `Body`       | <pre>interface{}</pre>  | Query body for this action               |
| `Index`      | <pre>string</pre>       | Index name for this action               |
| `Collection` | <pre>string</pre>       | Collection name for this action          |
| `Id`         | <pre>string</pre>       | id for this action                       |
| `Volatile`   | <pre>VolatileData</pre> | Additional information to send to Kuzzle |

### **options**

A [QueryOptions](https://github.com/kuzzleio/sdk-go/blob/master/types/query_options.go) containing additional query options
Theses properties can bet Get/Set.
The following properties are the most common.

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | Make this request queuable or not | true    |

### **responseChannel**

A channel to receive the API response.
This channel will receive a [KuzzleResponse](https://github.com/kuzzleio/sdk-go/blob/master/types/kuzzle_response.go)

## Usage

<<< ./snippets/query.go
