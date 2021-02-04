jwt, _ := kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)

res, err := kuzzle.Auth.CheckToken(jwt)

if err != nil {
  log.Fatal(err)
} else {
  if res.Valid != true {
    log.Fatal("Invalid token")
  }
  fmt.Println("Success")
}
