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
func kuzzle_wrapper_collection_count(c *C.collection, result *C.int_response, filters *C.filters, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).Count(filters, opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.result = C.int(res)
}
