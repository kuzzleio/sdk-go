package collection

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/kuzzleio/sdk-go/types"
//)
//
//// ScrollSpecifications retrieves next result of a specification search with scroll query.
//func (dc *Collection) ScrollSpecifications(scrollId string, options types.QueryOptions) (*types.SpecificationSearchResult, error) {
//	if scrollId == "" {
//		return nil, types.NewError("Collection.ScrollSpecifications: scroll id required")
//	}
//
//	ch := make(chan *types.KuzzleResponse)
//
//	query := &types.KuzzleRequest{
//		Controller: "collection",
//		Action:     "scrollSpecifications",
//		ScrollId:   scrollId,
//	}
//
//	if options != nil {
//		scroll := options.Scroll()
//		if scroll != "" {
//			query.Scroll = scroll
//		}
//	}
//
//	go dc.Kuzzle.Query(query, options, ch)
//
//	res := <-ch
//
//	if res.Error != nil {
//		return nil, res.Error
//	}
//
//	specifications := &types.SpecificationSearchResult{}
//	json.Unmarshal(res.Result, specifications)
//
//	return specifications, nil
//}
//
//// DeleteSpecifications deletes the current specifications of this collection.
//func (dc *Collection) DeleteSpecifications(options types.QueryOptions) (bool, error) {
//	ch := make(chan *types.KuzzleResponse)
//
//	query := &types.KuzzleRequest{
//		Collection: dc.collection,
//		Index:      dc.index,
//		Controller: "collection",
//		Action:     "deleteSpecifications",
//	}
//	go dc.Kuzzle.Query(query, options, ch)
//
//	res := <-ch
//
//	if res.Error != nil {
//		return false, res.Error
//	}
//
//	ack := &struct {
//		Acknowledged bool `json:"acknowledged"`
//	}{}
//	err := json.Unmarshal(res.Result, ack)
//	if err != nil {
//		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
//	}
//	return ack.Acknowledged, nil
//}
