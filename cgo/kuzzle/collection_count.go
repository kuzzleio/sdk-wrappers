package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_count
func kuzzle_wrapper_collection_count(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.int_result {
	result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.Count(cToGoSearchFilters(searchFilters), opts)

	if err != nil {
		Set_int_result_error(result, err)
		return result
	}

	result.result = C.int(res)

	return result
}