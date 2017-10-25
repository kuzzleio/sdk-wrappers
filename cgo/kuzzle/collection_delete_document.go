package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_delete_document
func kuzzle_wrapper_collection_delete_document(c *C.collection, id *C.char, options *C.query_options) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).DeleteDocument(C.GoString(id), opts)

	if err != nil {
		Set_string_result_error(result, err)
		return result
	}

	result.result = C.CString(res)

	return result
}
