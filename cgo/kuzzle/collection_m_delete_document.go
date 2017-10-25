package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"unsafe"
)

//export kuzzle_wrapper_collection_m_delete_document
func kuzzle_wrapper_collection_m_delete_document(c *C.collection, ids **C.char, idsCount C.uint, options *C.query_options) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))
	opts := SetQueryOptions(options)
	gIds := cToGoStrings(ids, idsCount)
	res, err := cToGoCollection(c).MDeleteDocument(gIds, opts)

	if err != nil {
		Set_string_array_result_error(result, err)
		return result
	}

	result.result = (**C.char)(C.calloc(C.size_t(len(res)), C.sizeof_char_ptr))
	result.length = C.ulong(len(res))

	cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(result.result))[:len(res):len(res)]

	for i, substring := range res {
		cArray[i] = C.CString(substring)
	}

	return result
}
