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
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_set_auto_refresh
func kuzzle_wrapper_set_auto_refresh(k *C.Kuzzle, index *C.char, auto_refresh C.uint, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	autoRefresh := auto_refresh != 0

	res, err := (*kuzzle.Kuzzle)(k.instance).SetAutoRefresh(C.GoString(index), autoRefresh, opts)
	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	if res {
		result.result = 1
	} else {
		result.result = 0
	}

	return result
}
