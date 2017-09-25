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

//export kuzzle_wrapper_collection_create_document
func kuzzle_wrapper_collection_create_document(c *C.collection, result *C.char, id *C.char, document *C.document, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).CreateDocument(C.GoString(id), *(*collection.Document)(document.instance), opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	*result = *(*[2048]C.char)(unsafe.Pointer(C.CString(res)))
}
