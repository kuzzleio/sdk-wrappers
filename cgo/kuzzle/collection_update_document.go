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

//export kuzzle_wrapper_collection_update_document
func kuzzle_wrapper_collection_update_document(c *C.collection, result *C.document, id *C.char, document *C.document, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).UpdateDocument(C.GoString(id), (*collection.Document)(document.instance), opts)
	if err != nil {
		if err.Error() == "Collection.UpdateDocument: document id required" {
			return C.int(C.EINVAL)
		}
		result.error = ToCString_2048(err.Error())
		return 0
	}

	result.instance = unsafe.Pointer(&res)

	return 0
}
