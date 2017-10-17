package types

type IKuzzle interface {
	Query(query *KuzzleRequest, options QueryOptions, responseChannel chan<- *KuzzleResponse)
}
