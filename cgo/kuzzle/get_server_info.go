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
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_get_server_info
func kuzzle_wrapper_get_server_info(k *C.Kuzzle, options *C.query_options) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetServerInfo(opts)

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
