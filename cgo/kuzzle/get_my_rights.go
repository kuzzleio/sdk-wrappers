package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include "kuzzle.h"
*/
import "C"
import (
	"unsafe"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_get_my_rights
func kuzzle_wrapper_get_my_rights(k *C.kuzzle, options *C.query_options) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
	result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
	opts := SetQueryOptions(options)

	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyRights(opts)

	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	r, _ := json.Marshal(res)
	buffer := C.CString(string(r))
	result.result.ptr = C.json_tokener_parse(buffer)

	C.free(unsafe.Pointer(buffer))
	return result
}
