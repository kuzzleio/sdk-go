isValid, err := kuzzle.Document.Validate(
	"nyc-open-data",
	"yellow-taxi",
	json.RawMessage(`{
  	"capacity": 4
	}`),
	nil)

if err != nil {
  log.Fatal(err)
} else if isValid {
  fmt.Println("Success")
}
