package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_collection_search
func kuzzle_wrapper_collection_search(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.search_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).Search(cToGoSearchFilters(searchFilters), opts)

	return goToCSearchResult(c, res, err)
}