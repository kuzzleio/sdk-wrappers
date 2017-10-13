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

//export kuzzle_wrapper_collection_count
func kuzzle_wrapper_collection_count(c *C.collection, result *C.int_response, searchFilters *C.search_filters, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}



	res, err := (*collection.Collection)(c.instance).Count(cToGoSearchFilters(searchFilters), opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.result = C.int(res)
}

func go_to_c_search_filter (cSf *C.search_filters) *types.SearchFilters {
	filters := types.SearchFilters{}

	if cSf.query != nil {
		filters.Query = JsonCConvert(cSf.query)
	}

	if cSf.sort != nil && JsonCType(cSf.sort) == C.json_type_array {
		filters.Sort = JsonCConvert(cSf.sort).([]interface{})
	}

	return &filters
}
