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
)

//export kuzzle_wrapper_update_self
func kuzzle_wrapper_update_self(k *C.kuzzle, credentials *C.json_object, options *C.query_options) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
	result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	jp := JsonParser{}
	jp.Parse(credentials)

	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateSelf(jp.GetContent(), opts)
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
