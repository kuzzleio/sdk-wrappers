package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
/*
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_search_specifications
// TODO
func kuzzle_wrapper_collection_search_specifications(c *C.collection, result *C.specification_search_result, searchFilters *C.search_filters, options *C.query_options) {
	opts := SetQueryOptions(options)
	col := cToGoCollection(c)
	res, err := col.SearchSpecifications(cToGoSearchFilters(searchFilters), opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	goToCSpecificationSearchResult(res, result)
}
*/
