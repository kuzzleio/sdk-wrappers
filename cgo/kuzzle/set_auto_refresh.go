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
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_set_auto_refresh
func kuzzle_wrapper_set_auto_refresh(k *C.kuzzle, result *C.bool_result, index *C.char, auto_refresh C.uint, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	var autoRefresh bool
	if auto_refresh == 1 {
		autoRefresh = true
	} else {
		autoRefresh = false
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).SetAutoRefresh(C.GoString(index), autoRefresh, opts)
	if err != nil {
		if err.Error() == "Kuzzle.SetAutoRefresh: index required" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
	}

	var r C.uint
	if res {
		r = 1
	} else {
		r = 0
	}

	result.result = r

	return 0
}
