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

//export kuzzle_wrapper_collection_document_exists
func kuzzle_wrapper_collection_document_exists(c *C.collection, result *C.bool_result, id *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).DocumentExists(C.GoString(id), opts)
	if err != nil {
		if err.Error() == "Collection.DocumentExists: document id required" {
			return C.int(C.EINVAL)
		} else {
			result.error = ToCString_2048(err.Error())
			return 0
		}
	}

	var valid C.uint

	if res {
		valid = 1
	} else {
		valid = 0
	}

	result.result = valid

	return 0
}
