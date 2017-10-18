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

//export kuzzle_wrapper_collection_delete_document
func kuzzle_wrapper_collection_delete_document(c *C.collection, result *C.string_result, id *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).DeleteDocument(C.GoString(id), opts)
	if err != nil {
		if err.Error() == "Collection.DeleteDocument: document id required" {
			return C.int(C.EINVAL)
		}
		result.error = ToCString_2048(err.Error())
		return 0
	}

	result.result = ToCString_2048(res)

	return 0
}
