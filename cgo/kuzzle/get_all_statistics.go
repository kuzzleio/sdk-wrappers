package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c

	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"unsafe"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(k *C.Kuzzle, options *C.query_options) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
	opts := SetQueryOptions(options)

	res, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)
	
	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	r, _ := json.Marshal(res)

	buffer := C.CString(string(r))
	defer C.free(unsafe.Pointer(buffer))

	result.result = C.json_tokener_parse(buffer)

	return result
}
