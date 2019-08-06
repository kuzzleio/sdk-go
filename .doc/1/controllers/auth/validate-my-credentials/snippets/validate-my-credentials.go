kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)
_, err := kuzzle.Auth.ValidateMyCredentials("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
