package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_now
func kuzzle_wrapper_now(k *C.kuzzle, res *C.now_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	time, err := (*kuzzle.Kuzzle)(k.instance).Now(opts)
	if err != nil {
		res.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
	}

	res.result = C.double(time)
}
