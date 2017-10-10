package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_collection_delete_document
func kuzzle_wrapper_collection_delete_document(c *C.collection, result *C.string_response, id *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).DeleteDocument(C.GoString(id), opts)
	if err != nil {
		if err.Error() == "Collection.DeleteDocument: document id required" {
			return C.int(C.EINVAL)
		}
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return 0
	}

	result.result = *(*[2048]C.char)(unsafe.Pointer(C.CString(res)))

	return 0
}