package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_search
func kuzzle_wrapper_collection_search(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)
	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.Search(cToGoSearchFilters(searchFilters), opts)

	return goToCSearchResult(res, c, err)
}