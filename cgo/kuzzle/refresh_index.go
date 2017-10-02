package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_refresh_index
func kuzzle_wrapper_refresh_index(k *C.Kuzzle, res *C.shards, index *C.char, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	shards, err := (*kuzzle.Kuzzle)(k.instance).RefreshIndex(C.GoString(index), opts)
	if err != nil {
		res.error = ToCString_2048(err.Error())
	}

	res.total = C.int(shards.Total)
	res.successful = C.int(shards.Successful)
	res.failed = C.int(shards.Failed)
}
