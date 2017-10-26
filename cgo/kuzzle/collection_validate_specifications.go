package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_validate_specifications
func kuzzle_wrapper_collection_validate_specifications(c *C.collection, specification *C.specification, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).ValidateSpecifications(cToGoSpecification(specification), opts)

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(res.Valid)

	return result
}
