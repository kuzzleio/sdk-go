err := kuzzle.Index.RefreshInternal(nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Internal index successfully refreshed")
}
