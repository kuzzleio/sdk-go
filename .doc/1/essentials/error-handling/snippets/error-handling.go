err := kuzzle.Index.Create("nyc-open-data", nil)

if err != nil {
  fmt.Println(err.Error())

  // Type assertion of error to KuzzleError
  if err.(types.KuzzleError).Status == 400 {
    fmt.Println("Try with another name!")
  }
}
