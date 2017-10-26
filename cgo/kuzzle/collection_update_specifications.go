package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_update_specifications
func kuzzle_wrapper_collection_update_specifications(c *C.collection, specification *C.specification, options *C.query_options) *C.specification_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).UpdateSpecifications(cToGoSpecification(specification), opts)

	return goToCSpecificationResult(res, err)
}
