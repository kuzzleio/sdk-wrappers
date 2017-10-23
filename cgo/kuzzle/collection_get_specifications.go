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
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_get_specifications
// TODO
func kuzzle_wrapper_collection_get_specifications(c *C.collection, result *C.kuzzle_specification, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.GetSpecifications(opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.instance = unsafe.Pointer(&res)
}
