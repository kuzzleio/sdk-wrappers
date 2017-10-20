package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_get_auto_refresh
func kuzzle_wrapper_get_auto_refresh(k *C.Kuzzle, index *C.char, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).GetAutoRefresh(C.GoString(index), opts)
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
