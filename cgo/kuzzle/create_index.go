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

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.ack_result {
	result := (*C.ack_result)(C.calloc(1, C.sizeof_ack_result))
	opts := SetQueryOptions(options)

	res, err := (*kuzzle.Kuzzle)(k.instance).CreateIndex(C.GoString(index), opts)

	if err != nil {
		Set_ack_result_error(result, err)
		return result
	}

	result.acknowledged = C.bool(res.Acknowledged)
	result.shards_acknowledged = C.bool(res.ShardsAcknowledged)

	return result
}
