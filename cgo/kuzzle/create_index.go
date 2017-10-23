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

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(k *C.Kuzzle, index *C.char, options *C.query_options) *C.ack_response {
	result := (*C.ack_response)(C.calloc(1, C.sizeof_ack_response))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).CreateIndex(C.GoString(index), opts)

	if err != nil {
		Set_ack_response_error(result, err)
		return result
	}

	if res.Acknowledged {
		result.acknowledged = 1
	} else {
		result.acknowledged = 0
	}

	if res.ShardsAcknowledged {
		result.shardsAcknowledged = 1
	} else {
		result.shardsAcknowledged = 0
	}

	return result
}
