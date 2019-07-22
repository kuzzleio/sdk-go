kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)
_, err := kuzzle.Auth.GetMyCredentials("local", nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
