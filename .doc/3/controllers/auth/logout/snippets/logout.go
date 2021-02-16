kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)
err := kuzzle.Auth.Logout()

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
