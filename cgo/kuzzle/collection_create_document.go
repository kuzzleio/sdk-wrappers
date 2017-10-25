package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_create_document
func kuzzle_wrapper_collection_create_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).CreateDocument(C.GoString(id), cToGoDocument(c, document), opts)

	if err != nil {
		Set_document_error(result, err)
		return result
	}

	result.result = goToCDocument(c, res)

	return result
}
