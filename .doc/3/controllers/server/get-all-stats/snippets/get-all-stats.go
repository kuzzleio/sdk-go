_, err := kuzzle.Server.GetAllStats(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Success")
}
