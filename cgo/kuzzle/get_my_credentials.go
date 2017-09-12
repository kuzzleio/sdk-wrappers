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

//export kuzzle_wrapper_get_my_credentials
func kuzzle_wrapper_get_my_credentials(k *C.kuzzle, result *C.json_result, strategy *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyCredentials(C.GoString(strategy), opts)
	if err != nil {
		if err.Error() == "Kuzzle.GetMyCredentials: strategy is required" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
	}

	r, _ := json.Marshal(res)
	result.result = C.json_tokener_parse(C.CString(string(r)))

	return 0
}
