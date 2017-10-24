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

//export kuzzle_wrapper_now
func kuzzle_wrapper_now(k *C.kuzzle, options *C.query_options) *C.int_result {
	result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	time, err := (*kuzzle.Kuzzle)(k.instance).Now(opts)
	if err != nil {
		Set_int_result_error(result, err)
		return result
	}

	result.result = C.longlong(time)
	return result
}
