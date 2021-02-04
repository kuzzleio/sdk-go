start := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
stop := time.Now()

stats, err := kuzzle.Server.GetStats(&start, &stop, nil)

if err != nil {
  log.Fatal(err)
} else {
  fmt.Println("Kuzzle Stats as JSON string:", string(stats))
}
