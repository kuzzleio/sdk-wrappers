package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_collection_search_specifications
func kuzzle_wrapper_collection_search_specifications(c *C.collection, result *C.kuzzle_specification_search_response, searchFilters *C.search_filters, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).SearchSpecifications(goToCSearchFilters(searchFilters), opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	gotToCSpecificationSearchResult(res, result)
}