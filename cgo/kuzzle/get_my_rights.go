package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <json/json.h>
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"encoding/json"
)

//export kuzzle_wrapper_get_my_rights
func kuzzle_wrapper_get_my_rights(result *C.json_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := KuzzleInstance.GetMyRights(opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	r, _ := json.Marshal(res)
	result.result = C.json_tokener_parse(C.CString(string(r)))
}
