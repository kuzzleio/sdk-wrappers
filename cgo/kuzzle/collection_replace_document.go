package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_collection_replace_document
func kuzzle_wrapper_collection_replace_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := cToGoCollection(c).ReplaceDocument(C.GoString(id), cToGoDocument(c, document), opts)

	if err != nil {
		Set_document_error(result, err)
		return result
	}

	result.result = goToCDocument(c, res)

	return result
}
