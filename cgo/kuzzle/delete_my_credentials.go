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

//export kuzzle_wrapper_delete_my_credentials
func kuzzle_wrapper_delete_my_credentials(k *C.Kuzzle, strategy *C.char, options *C.query_options) *C.ack_response {
	result := (*C.ack_response)(C.calloc(1, C.sizeof_ack_response))
	opts := SetQueryOptions(options)

	res, err := (*kuzzle.Kuzzle)(k.instance).DeleteMyCredentials(C.GoString(strategy), opts)

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
