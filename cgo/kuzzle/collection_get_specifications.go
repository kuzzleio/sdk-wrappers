package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_get_specifications
func kuzzle_wrapper_collection_get_specifications(c *C.collection, options *C.query_options) *C.specification_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).GetSpecifications(opts)

	return goToCSpecificationResult(res.Validation, err)
}
