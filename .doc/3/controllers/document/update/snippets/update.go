kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-id",
	json.RawMessage(`{"capacity": 4}`),
	nil)

response, err := kuzzle.Document.Update(
	"nyc-open-data",
	"yellow-taxi",
	"some-id",
	json.RawMessage(`{"category": "suv"}`),
	nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println(string(response))
  /*
  {
    "_index": "nyc-open-data",
    "_type": "yellow-taxi",
    "_id": "some-id",
    "_version": 2,
    "result": "updated",
    "_shards": {
      "total": 2,
      "successful": 1,
      "failed": 0
    }
  }
  */
}
