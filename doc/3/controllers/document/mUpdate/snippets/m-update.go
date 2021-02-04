kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-id",
	json.RawMessage(`{"capacity": 4}`),
	nil);

kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-other-id",
	json.RawMessage(`{"capacity": 7}`),
	nil);

response, err := kuzzle.Document.MReplace(
	"nyc-open-data",
	"yellow-taxi",
	json.RawMessage(`[
		{
			"_id": "some-id",
			"body": { "category": "sedan" }
		},
		{
			"_id": "some-other-id",
			"body": { "category": "limousine" }
		}
	]`),
	nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println(string(response))
  fmt.Println("Success")
}