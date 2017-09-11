package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_validate_my_credentials
func kuzzle_wrapper_validate_my_credentials(result *C.bool_result, strategy *C.char, credentials *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	jp := JsonParser{}
	jp.Parse(credentials)

	res, err := KuzzleInstance.ValidateMyCredentials(C.GoString(strategy), jp.GetContent(), opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
	}

	var r C.uint
	if res {
		r = 1
	} else {
		r = 0
	}
	result.result = r
}
