package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <stdlib.h>
	#include "kuzzle.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_list_indexes
func kuzzle_wrapper_list_indexes(k *C.kuzzle, options *C.query_options) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).ListIndexes(opts)
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
