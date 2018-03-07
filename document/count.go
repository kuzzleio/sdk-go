package document

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) Count(index string, collection string, body string) (int, error) {
	if index == "" {
		return 0, types.NewError("Document.Count: index required", 400)
	}

	if collection == "" {
		return 0, types.NewError("Document.Count: collection required", 400)
	}

	if body == "" {
		return 0, types.NewError("Document.Count: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "count",
		Body:       body,
	}
	go d.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return 0, res.Error
	}

	var count int
	err := json.Unmarshal(res.Result, &count)

	if err != nil {
		return 0, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return count, nil
}
