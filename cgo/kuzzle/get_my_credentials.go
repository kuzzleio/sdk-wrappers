package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_get_my_credentials
func kuzzle_wrapper_get_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
	result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyCredentials(C.GoString(strategy), opts)

	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	r, _ := json.Marshal(res)
	buffer := C.CString(string(r))
	defer C.free(unsafe.Pointer(buffer))

	result.result.ptr = C.json_tokener_parse(buffer)

	return result
}
