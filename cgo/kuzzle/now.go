package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_now
func kuzzle_wrapper_now(k *C.Kuzzle, res *C.now_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	time, err := (*kuzzle.Kuzzle)(k.instance).Now(opts)
	if err != nil {
		res.error = ToCString_2048(err.Error())
	}

	res.result = C.double(time)
}
