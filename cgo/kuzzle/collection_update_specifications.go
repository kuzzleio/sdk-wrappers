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
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_update_specifications
// TODO
func kuzzle_wrapper_collection_update_specifications(c *C.collection, specification *C.kuzzle_specification, result *C.kuzzle_specification, options *C.query_options) {
	opts := SetQueryOptions(options)
	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.UpdateSpecifications((*types.KuzzleValidation)(specification.instance), opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.result = unsafe.Pointer(res)
}
*/