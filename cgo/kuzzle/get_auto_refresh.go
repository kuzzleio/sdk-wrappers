package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_get_auto_refresh
func kuzzle_wrapper_get_auto_refresh(k *C.Kuzzle, result *C.bool_result, index *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetAutoRefresh(C.GoString(index), opts)
	if err != nil {
		if err.Error() == "Kuzzle.CheckToken: token required" {
			return C.int(C.EINVAL)
		} else {
			result.error = ToCString_2048(err.Error())
			return 0
		}
	}

	var valid C.uint

	if res {
		valid = 1
	} else {
		valid = 0
	}

	result.result = valid

	return 0
}
