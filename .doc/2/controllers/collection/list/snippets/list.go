options := types.NewQueryOptions()
options.SetFrom(1)
options.SetSize(1)

list, err := kuzzle.Collection.List("mtp-open-data", options)

if err != nil {
  log.Fatal(err)
} else if list != nil {
  fmt.Println("Success")
}
