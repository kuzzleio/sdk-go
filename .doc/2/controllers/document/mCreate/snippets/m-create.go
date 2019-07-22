documents := json.RawMessage(`[
  {
    "_id": "some-id",
    "body": { "capacity": 4 }
  },
  {
    "body": { "this": "document id is auto-computed" }
  }
]`)
response, err := kuzzle.Document.MCreate(
	"nyc-open-data",
	"yellow-taxi",
	documents,
	nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println(string(response))
  /*
  [
    {
      "_id":"some-id",
      "_source":{
          "_kuzzle_info":{
            "active":true,
            "author":"-1",
            "updater":null,
            "updatedAt":null,
            "deletedAt":null,
            "createdAt":1538484279484
          },
          "capacity":4
      },
      "_index":"nyc-open-data",
      "_type":"yellow-taxi",
      "_version":1,
      "result":"created",
      "_shards":{
          "total":2,
          "successful":1,
          "failed":0
      },
      "created":true,
      "status":201
    },
    {
      "_id":"AWY0zxi_7XvER2v0e9xR",
      "_source":{
          "_kuzzle_info":{
            "active":true,
            "author":"-1",
            "updater":null,
            "updatedAt":null,
            "deletedAt":null,
            "createdAt":1538484279484
          },
          "this":"document id is auto-computed"
      },
      "_index":"nyc-open-data",
      "_type":"yellow-taxi",
      "_version":1,
      "result":"created",
      "_shards":{
          "total":2,
          "successful":1,
          "failed":0
      },
      "created":true,
      "status":201
    }
  ]
  */
  fmt.Println("Success")
}