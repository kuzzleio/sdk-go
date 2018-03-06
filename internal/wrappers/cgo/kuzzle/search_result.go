package main

/*
  #cgo CFLAGS: -I../../headers
  #include "kuzzlesdk.h"
*/
import "C"

////export kuzzle_search_result_fetch_next
//func kuzzle_search_result_fetch_next(sr *C.search_result) *C.search_result {
//	res, err := cToGoSearchResult(sr).FetchNext()
//	return goToCSearchResult(sr.collection, res, err)
//}
