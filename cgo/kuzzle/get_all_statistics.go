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
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(k *C.kuzzle, result *C.json_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	r, _ := json.Marshal(res)
	result.result = C.json_tokener_parse(C.CString(string(r)))
}
