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

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(k *C.Kuzzle, result *C.json_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	r, _ := json.Marshal(res)
	cString := C.CString(string(r))
	defer C.free(unsafe.Pointer(cString))
	result.result = C.json_tokener_parse(cString)
}
