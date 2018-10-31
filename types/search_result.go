// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchResult is a search result representation
type SearchResult struct {
	Aggregations json.RawMessage
	Hits         json.RawMessage
	Total        int
	Fetched      int
	ScrollId     string
	kuzzle       IKuzzle
	request      *KuzzleRequest
	response     *KuzzleResponse
	options      QueryOptions
	scrollAction string
}

// NewSearchResult Search Result constructor
func NewSearchResult(kuzzle IKuzzle, scrollAction string, request *KuzzleRequest, options QueryOptions, raw *KuzzleResponse) (*SearchResult, error) {
	type ParseSearchResult struct {
		Hits         json.RawMessage `json:"hits"`
		Total        int             `json:"total"`
		ScrollId     string          `json:"_scroll_id"`
		Aggregations json.RawMessage `json:"aggregations"`
	}

	var parsed ParseSearchResult
	err := json.Unmarshal(raw.Result, &parsed)

	if err != nil {
		return nil, NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), raw.Result), 500)
	}

	var docs []interface{}
	json.Unmarshal(parsed.Hits, &docs)

	sr := &SearchResult{
		Aggregations: parsed.Aggregations,
		Hits:         parsed.Hits,
		Fetched:      len(docs),
		Total:        parsed.Total,
		ScrollId:     parsed.ScrollId,
		kuzzle:       kuzzle,
		request:      request,
		response:     raw,
		options:      NewQueryOptions(),
		scrollAction: scrollAction,
	}

	sr.options = options

	return sr, nil
}

func (sr *SearchResult) Kuzzle() IKuzzle {
	return sr.kuzzle
}

func (sr *SearchResult) Request() *KuzzleRequest {
	return sr.request
}

func (sr *SearchResult) Response() *KuzzleResponse {
	return sr.response
}

func (sr *SearchResult) Options() QueryOptions {
	return sr.options
}

func (sr *SearchResult) ScrollAction() string {
	return sr.scrollAction
}

// Next page result
func (sr *SearchResult) Next() (*SearchResult, error) {
	if sr.Fetched >= sr.Total {
		return nil, nil
	}

	var pb struct {
		Sort []map[string]string `json:"sort"`
	}
	j, _ := json.Marshal(sr.request.Body)
	json.Unmarshal(j, &pb)

	if sr.ScrollId != "" {
		// scroll

		query := &KuzzleRequest{
			Controller: sr.request.Controller,
			Action:     sr.scrollAction,
		}

		options := sr.options
		options.SetScrollId(sr.ScrollId)

		ch := make(chan *KuzzleResponse)
		go sr.kuzzle.Query(query, sr.options, ch)
		res := <-ch

		nsr, err := NewSearchResult(sr.kuzzle, sr.scrollAction, query, options, res)

		if err != nil {
			return nil, err
		}

		nsr.Fetched += sr.Fetched

		return nsr, nil
	} else if pb.Sort != nil && sr.request.Size > 0 {
		// search_after

		query := sr.request
		query.RequestId = ""

		type Parsed struct {
			Hits []map[string]interface{} `json:"hits"`
		}
		var parsed Parsed
		json.Unmarshal(sr.response.Result, &parsed)

		last := make(map[string]interface{})
		if len(parsed.Hits) > 0 {
			last = parsed.Hits[len(parsed.Hits)-1]
		}

		searchAfter := make([]interface{}, 0)
		for _, v := range pb.Sort {
			for key := range v {
				searchAfter = append(searchAfter, getDeepField(key, last, 0))
			}
		}

		var body map[string]interface{}
		jbody, ok := sr.request.Body.(json.RawMessage)
		if ok {
			json.Unmarshal(jbody, &body)
			body["search_after"] = searchAfter
			query.Body = body
		}
		mbody, ok := sr.request.Body.(map[string]interface{})
		if ok {
			mbody["search_after"] = searchAfter
			query.Body = mbody
		}
		fmt.Println(searchAfter)

		options := sr.options

		ch := make(chan *KuzzleResponse)
		go sr.kuzzle.Query(query, sr.options, ch)
		res := <-ch

		nsr, err := NewSearchResult(sr.kuzzle, sr.scrollAction, query, options, res)

		if err != nil {
			return nil, err
		}

		nsr.Fetched += sr.Fetched

		return nsr, nil
	} else if sr.request.Size > 0 {
		// from/size

		if sr.request.From >= sr.Total {
			return nil, nil
		}

		query := sr.request
		query.RequestId = ""
		query.From = sr.Fetched
		options := sr.options
		options.SetFrom(sr.Fetched)

		ch := make(chan *KuzzleResponse)
		go sr.kuzzle.Query(query, options, ch)
		res := <-ch

		nsr, err := NewSearchResult(sr.kuzzle, sr.scrollAction, query, options, res)

		if err != nil {
			return nil, err
		}

		nsr.Fetched += sr.Fetched

		return nsr, nil
	}

	return nil, NewError("Unable to retrieve next results from search: missing scrollId, from/sort, or from/size parameters")
}

func getDeepField(key string, s map[string]interface{}, depth int) interface{} {
	if depth == 0 && key == "_uid" {
		key = "_id"
	}
	if depth == 0 && key != "_id" && key != "_score" {
		key = "_source." + key
	}

	if s[key] != nil {
		str, ok := s[key].(string)
		if ok {
			return str
		}
		i, ok := s[key].(int)
		if ok {
			return i
		}
		f, ok := s[key].(float64)
		if ok {
			return f
		}
	}

	keys := strings.Split(key, ".")
	if len(keys) == 1 {
		return nil
	}

	m, ok := s[keys[0]].(map[string]interface{})
	if ok {
		return getDeepField(strings.Join(keys[1:], "."), m, depth+1)
	}

	return nil
}
