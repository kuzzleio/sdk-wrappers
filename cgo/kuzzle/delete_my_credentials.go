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

//export kuzzle_wrapper_delete_my_credentials
func kuzzle_wrapper_delete_my_credentials(k *C.kuzzle, result *C.ack_response, strategy *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).DeleteMyCredentials(C.GoString(strategy), opts)
	if err != nil {
		if err.Error() == "Kuzzle.DeleteMyCredentials: strategy is required" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
	}

	var ack, shardsAck C.uint

	if res.Acknowledged {
		ack = 1
	} else {
		ack = 0
	}

	if res.ShardsAcknowledged {
		shardsAck = 1
	} else {
		shardsAck = 0
	}

	*result = C.ack_response{
		acknowledged:       ack,
		shardsAcknowledged: shardsAck,
	}

	return 0
}
