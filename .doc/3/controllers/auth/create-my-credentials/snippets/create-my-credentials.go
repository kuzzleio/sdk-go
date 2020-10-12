kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)
kuzzle.Auth.CreateMyCredentials("other", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)

fmt.Println("Success")