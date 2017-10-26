package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_refresh_index
func kuzzle_wrapper_refresh_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.shards_result {
	result := (*C.shards_result)(C.calloc(1, C.sizeof_shards_result))
	opts := SetQueryOptions(options)

	shards, err := (*kuzzle.Kuzzle)(k.instance).RefreshIndex(C.GoString(index), opts)
	if err != nil {
		Set_shards_result_error(result, err)
		return result
	}

	result.result = goToCShards(shards)

	return result
}
