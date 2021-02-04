kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)
_, err := kuzzle.Auth.UpdateMyCredentials("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\",\"other\":\"value\"}"), nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
