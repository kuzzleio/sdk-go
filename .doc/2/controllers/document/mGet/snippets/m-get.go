kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-id",
	json.RawMessage(`{}`),
	nil)

kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-other-id",
	json.RawMessage(`{}`),
	nil)

ids := []string{"some-id", "some-other-id"}
response, err := kuzzle.Document.MGet("nyc-open-data", "yellow-taxi", ids, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println(string(response))
  fmt.Println("Success")
}
