package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"unsafe"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/collection"
)

//export kuzzle_wrapper_collection_create
func kuzzle_wrapper_collection_create(c *C.collection, result *C.toto, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).Create(opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	res, err := KuzzleInstance.CheckToken(C.GoString(token))
	if err != nil {
		if err.Error() == "Kuzzle.CheckToken: token required" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
	}

	var valid C.uint

	if res.Valid {
		valid = 1
	} else {
		valid = 0
	}

	*result = C.token_validity{
		valid:     valid,
		state:     *(*[512]C.char)(unsafe.Pointer(C.CString(res.State))),
		expiresAt: C.int(res.ExpiresAt),
	}

	return 0
}
