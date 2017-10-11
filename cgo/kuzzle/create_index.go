package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(k *C.Kuzzle, result *C.ack_response, index *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).CreateIndex(C.GoString(index), opts)
	if err != nil {
		if err.Error() == "Collection.createIndex: index required" {
			return C.int(C.EINVAL)
		} else {
			result.error = C.CString(err.Error())
			return 0
		}
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

	return 0
}
