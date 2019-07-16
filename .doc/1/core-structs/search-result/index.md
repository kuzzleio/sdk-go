---
code: true
type: page
title: SearchResult
description: How to deal with Kuzzle search results
order: 600
---

# SearchResult

When performing a [search](/sdk/go/1/controllers/document/search/), Kuzzle returns an `SearchResult` struct, which holds the items matching the given query and allows to drill through next result pages if applicable.

## Properties

| Property       | Type                       | Description                                                           |
| -------------- | -------------------------- | --------------------------------------------------------------------- |
| `Aggregations` | <pre>json.RawMessage</pre> | The search` aggregations if any                                       |
| `Hits`         | <pre>json.RawMessage</pre> | A JSON string containing the retrieved items for the current page     |
| `Total`        | <pre>int</pre>             | The total number of items matching the given query in Kuzzle database |
| `Fetched`      | <pre>int</pre>             | The number of retrieved items so far                                  |
| `ScrollId`     | <pre>string</pre>          | A scroll identifier if the search was given a `scroll` parameter      |

## Methods

The `SearchResult` struct exposes an unique `Next` method, which returns a new `SearchResult` struct.

```cpp
Next() (*SearchResult, error)
```

## Behavior of the next method

In order to be able to compute the next search page, some initial conditions must be met.

Depending on the arguments given to the initial search, thhe `Next` method will pick one of the following policies, by decreasing order of priority (i.e. a search including a `scroll`, `sort` and `size` will use the `scroll` policy).

If no policy is applicable, the `next` method will throw an exception.

:::info
When processing a large number of documents (i.e. more than 1000), it is advised to use a scroll cursor.

It is also the only method that garantees all matching documents will be retrieved and no duplicates will be included.
:::

### 1. scroll

**This is the preferred way to get some paginated results**.

If the original search is given a `scroll` parameter, the `next` method will use a cursor to paginate results.

The results that are returned from a scroll request reflect the state of the index at the time the initial `search` request was performed, like a snapshot in time.

As such, even if some documents are added or deleted from the database between two calls to `Next`, the result is garanteed to include all items matching the query at the time the initial `search` was sent and to not get any duplicate between two search pages.

### 2. sort / size

If the initial search is given some `sort` and `size` parameters, the `Next` method will retrieve the next items matching the sort.

To avoid too many duplicates, it is advised to provide a sort combination that will always identify one item only. The recommended way is to use the field `_uid` which is certain to contain one unique value for each document.

Because this method does not freeze the research between two calls, if some updates are applied to the database between two calls, it is still possible to miss some documents and/or to get some duplicates between search pages.

### 3. from / size

If the initial search is given some `from` and `size` parameters, the `Next` method will increment the `from` parameter to retrieved the next results.

Because this method does not freeze the research between two calls, if some updates are applied to the database between two calls, it is possible to miss some documents and/or to get some duplicates between search pages.

:::info
It is not possible to retrieve more than 10000 items using this method. Above that limit, any call to `Next` will return an error
:::

## Usage

<<< ./snippets/search-result.go
