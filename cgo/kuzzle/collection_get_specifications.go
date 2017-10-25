package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
/*
import (
	"github.com/kuzzleio/sdk-go/collection"
	"unsafe"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_get_specifications
// TODO
func kuzzle_wrapper_collection_get_specifications(c *C.collection, result *C.specification, options *C.query_options) {
	opts := SetQueryOptions(options)

	col := cToGoCollection(c)
	res, err := col.GetSpecifications(opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.instance = unsafe.Pointer(&res)
}
*/