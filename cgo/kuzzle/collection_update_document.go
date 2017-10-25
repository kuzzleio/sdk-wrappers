package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_update_document
func kuzzle_wrapper_collection_update_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))
	opts := SetQueryOptions(options)
	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.UpdateDocument(C.GoString(id), cToGoDocument(document, c), opts)

	if err != nil {
		Set_document_error(result, err)
		return result
	}

	result.result = goToCDocument(res, c)

	return result
}
