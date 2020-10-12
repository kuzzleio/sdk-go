kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-id",
	json.RawMessage("{}"),
	nil);

kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-other-id",
	json.RawMessage("{}"),
	nil);

response, err := kuzzle.Document.MReplace(
	"nyc-open-data",
	"yellow-taxi",
	json.RawMessage(`[
		{
			"_id": "some-id",
			"body": { "capacity": 4 }
		},
		{
			"_id": "some-other-id",
			"body": { "capacity": 4 }
		}
	]`),
	nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println(string(response))
  fmt.Println("Success")
}
