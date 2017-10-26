package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_search_specifications
func kuzzle_wrapper_collection_search_specifications(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.specification_search_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).SearchSpecifications(cToGoSearchFilters(searchFilters), opts)

	return goToCSpecificationSearchResult(res, err)
}

