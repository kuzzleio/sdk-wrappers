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

//export kuzzle_wrapper_list_collections
func kuzzle_wrapper_list_collections(result *C.json_result, index *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := KuzzleInstance.ListCollections(C.GoString(index), opts)
	if err != nil {
		if err.Error() == "Kuzzle.ListCollections: index required" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
	}

	var jsonRes *C.json_object
	r, _ := json.Marshal(res)

	jsonRes = C.json_tokener_parse(C.CString(string(r)))
	result.result = jsonRes

	return 0
}
