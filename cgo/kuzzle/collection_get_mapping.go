package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_collection_get_mapping
func kuzzle_wrapper_collection_get_mapping(c *C.collection, result *C.collection_mapping, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).GetMapping(opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.instance = unsafe.Pointer(&res)
}
