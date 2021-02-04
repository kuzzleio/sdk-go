_, err := kuzzle.Document.Create(
	"nyc-open-data",
	"yellow-taxi",
	"some-id",
	json.RawMessage(`{"capacity": 4}`),
	nil)

response, err := kuzzle.Document.Get("nyc-open-data", "yellow-taxi", "some-id", nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println(string(response))
  /*
  {
    "_index":"nyc-open-data",
    "_type":"yellow-taxi",
    "_id":"some-id",
    "_version":1,
    "found":true,
    "_source":{
      "capacity":4,
      "_kuzzle_info":{
          "author":"-1",
          "createdAt":1538402859880,
          "updatedAt":null,
          "updater":null,
          "active":true,
          "deletedAt":null
      }
    }
  }
  */
  fmt.Println("Success")
}
