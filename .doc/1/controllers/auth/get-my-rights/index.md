---
code: true
type: page
title: GetMyRights
description: Returns the rights for the user linked to the `JSON Web Token`.
---

# GetMyRights

Returns the rights for the user linked to the `JSON Web Token`, provided in the query or the `Authorization` header.

## Arguments

```go
func (a *Auth) GetMyRights(options types.QueryOptions) ([]*types.UserRights, error)
```

| Arguments | Type         | Description                                  | Required |
| --------- | ------------ | -------------------------------------------- | -------- |
| `options` | <pre>QueryOptions</pre> | QueryOptions object containing query options | yes      |

### **Options**

Additional query options

| Property   | Type | Description                       | Default |
| ---------- | ---- | --------------------------------- | ------- |
| `Queuable` | <pre>bool</pre> | Make this request queuable or not | `true`  |

## Return

A pointer to an array of UserRight object containing:

| Property      | Type   | Description                               |
| ------------- | ------ | ----------------------------------------- |
| `Controller`  | <pre>string</pre> | Controller on wich the rights are applied |
| `Action`      | <pre>string</pre> | Action on wich the rights are applied     |
| `Index`       | <pre>string</pre> | Index on wich the rights are applied      |
|  `Collection` | <pre>string</pre> | Collection on wich the rights are applied |
|  `Value`      | <pre>string</pre> | Rights (`allowed|denied|conditional`)     |

and an error or `nil`

## Usage

<<< ./snippets/get-my-rights.go
