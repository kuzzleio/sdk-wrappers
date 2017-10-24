package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_validate_specifications
// TODO
func kuzzle_wrapper_collection_validate_specifications(c *C.collection, specification *C.kuzzle_specification, result *C.bool_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.ValidateSpecifications((*types.KuzzleValidation)(specification.instance), opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	var valid C.uint

	if res.Valid {
		valid = 1
	} else {
		valid = 0
	}

	result.result = valid
}