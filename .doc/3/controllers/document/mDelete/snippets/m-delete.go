ids := []string{"some-id", "some-other-id"}

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

	
deletedJSON, err := kuzzle.Document.MDelete("nyc-open-data", "yellow-taxi", ids, nil)

if err != nil {
  log.Fatal(err)
} else {
	type deletedResult struct {
		Successes []json.RawMessage `json:"successes"`
		Errors 	  []json.RawMessage `json:"errors"`
	}
	var deleted deletedResult
	json.Unmarshal(deletedJSON, &deleted)
  	fmt.Printf("Successfully deleted %d documents", len(deleted.Successes))
}
