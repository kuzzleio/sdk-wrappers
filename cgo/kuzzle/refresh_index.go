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

//export kuzzle_wrapper_refresh_index
func kuzzle_wrapper_refresh_index(k *C.kuzzle, res *C.shards, index *C.char, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	shards, err := (*kuzzle.Kuzzle)(k.instance).RefreshIndex(C.GoString(index), opts)
	if err != nil {
		res.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
	}

	res.total = C.int(shards.Total)
	res.successful = C.int(shards.Successful)
	res.failed = C.int(shards.Failed)
}
