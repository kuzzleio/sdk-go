---
code: true
type: page
title: Login
description: Authenticates a user
---

# Login

Authenticates a user.

If this action is successful, all further requests emitted by this SDK instance will be in the name of the authenticated user, until either the authenticated token expires, the [logout](/sdk/go/1/controllers/auth/logout) action is called, or the `jwt` property is manually unset.

## Arguments

```go
func (a *Auth) Login(
  strategy string,
  credentials json.RawMessage,
  expiresIn *int) (string, error)
```


| Arguments     | Type   | Description                      | Required |
| ------------- | ------ | -------------------------------- | -------- |
| `strategy`    | <pre>string</pre> | Name of the strategy to use  | yes      |
| `credentials` | <pre>string</pre> |  Credentials for that strategy            |  yes     |
| `expiresIn`   | <pre>int</pre>    |  Expiration time, in milliseconds |  no      |

#### strategy

The name of the authentication [strategy](/core/1/guides/kuzzle-depth/authentication/#authentication) used to log the user in.

Depending on the chosen authentication `strategy`, additional [credential arguments](/core/1/guides/kuzzle-depth/authentication/#authentication) may be required.
The API request example in this page provides the necessary arguments for the [`local` authentication plugin](https://github.com/kuzzleio/kuzzle-plugin-auth-passport-local).

Check the appropriate [authentication plugin](/core/1/plugins/guides/strategies/overview/) documentation to get the list of additional arguments to provide.

### expiresIn
 The default value for the `expiresIn` option is defined at server level, in Kuzzle's [configuration file](/core/1/guides/essentials/configuration/).


## Return

The **login** action returns an encrypted JSON Web Token, that must then be sent in the [requests headers](/core/1/api/essentials/query-syntax/).

## Usage

<<< ./snippets/login.go
