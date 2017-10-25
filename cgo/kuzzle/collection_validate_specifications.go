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
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_validate_specifications
// TODO
func kuzzle_wrapper_collection_validate_specifications(c *C.collection, specification *C.kuzzle_specification, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
	opts := SetQueryOptions(options)
	col := cToGoCollection(c)
	res, err := col.ValidateSpecifications((*types.KuzzleValidation)(specification.instance), opts)

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(res)

	return result
}
*/