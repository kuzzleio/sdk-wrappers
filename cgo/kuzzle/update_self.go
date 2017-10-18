package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_update_self
func kuzzle_wrapper_update_self(k *C.Kuzzle, result *C.json_result, credentials *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateSelf(JsonCConvert(credentials).(map[string]interface{}), opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
	}

	var jsonRes *C.json_object
	r, _ := json.Marshal(res)

	cString := C.CString(string(r))
	defer C.free(unsafe.Pointer(cString))
	jsonRes = C.json_tokener_parse(cString)
	result.result = jsonRes
}
