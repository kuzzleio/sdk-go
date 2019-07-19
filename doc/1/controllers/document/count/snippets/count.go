query := json.RawMessage(`{"query": {"match": {"licence": "valid"}}}`)
count, err := kuzzle.Document.Count("nyc-open-data", "yellow-taxi", query, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Printf("Found %d documents matching licence:valid\n", count)
}
