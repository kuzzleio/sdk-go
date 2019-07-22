kuzzle.Auth.Login("local", json.RawMessage("{\"username\":\"foo\",\"password\":\"bar\"}"), nil)
res, err := kuzzle.Auth.CredentialsExist("local", nil)

if err != nil {
  log.Fatal(err)
} else {
	if res == true {
		fmt.Println("Success")
	} else {
	 log.Fatal("Error")
  }
}
