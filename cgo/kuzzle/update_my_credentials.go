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

//export kuzzle_wrapper_update_my_credentials
func kuzzle_wrapper_update_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
	result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
	opts := SetQueryOptions(options)

	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateMyCredentials(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), opts)
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
