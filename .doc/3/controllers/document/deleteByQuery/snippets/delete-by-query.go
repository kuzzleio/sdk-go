ids, err := kuzzle.Document.DeleteByQuery(
	"nyc-open-data",
	"yellow-taxi",
	json.RawMessage(`{"query": {"term": {"capacity": 7}}}`),
	nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Successfully deleted %d documents", len(ids))
}