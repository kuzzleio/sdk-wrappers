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

//export kuzzle_wrapper_collection_delete_document
func kuzzle_wrapper_collection_delete_document(c *C.collection, id *C.char, options *C.query_options) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))
	opts := SetQueryOptions(options)

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.DeleteDocument(C.GoString(id), opts)

	if err != nil {
		Set_string_result_error(result, err)
		return result
	}

	result.result = C.CString(res)

	return result
}
