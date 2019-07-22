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

deleted, err := kuzzle.Document.MDelete("nyc-open-data", "yellow-taxi", ids, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Successfully deleted %d documents", len(deleted))
}
