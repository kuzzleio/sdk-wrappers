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

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(result *C.json_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := KuzzleInstance.GetAllStatistics(opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	r, _ := json.Marshal(res)
	result.result = C.json_tokener_parse(C.CString(string(r)))
}
