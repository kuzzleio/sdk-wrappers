package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"unsafe"
)

//export kuzzle_wrapper_update_self
func kuzzle_wrapper_update_self(result *C.json_result, credentials *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	jp := JsonParser{}
	jp.Parse(credentials)

	res, err := KuzzleInstance.UpdateSelf(jp.GetContent(), opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
	}

	var jsonRes *C.json_object
	r, _ := json.Marshal(res)

	jsonRes = C.json_tokener_parse(C.CString(string(r)))
	result.result = jsonRes
}
