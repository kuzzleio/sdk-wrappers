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

//export kuzzle_wrapper_list_collections
func kuzzle_wrapper_list_collections(k *C.Kuzzle, result *C.json_result, index *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).ListCollections(C.GoString(index), opts)
	if err != nil {
		if err.Error() == "Kuzzle.ListCollections: index required" {
			return C.int(C.EINVAL)
		} else {
			result.error = ToCString_2048(err.Error())
			return 0
		}
	}

	var jsonRes *C.json_object
	r, _ := json.Marshal(res)

	cString := C.CString(string(r))
	defer C.free(unsafe.Pointer(cString))
	jsonRes = C.json_tokener_parse(cString)
	result.result = jsonRes

	return 0
}
