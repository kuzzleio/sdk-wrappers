package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_get_mapping
// TODO
func kuzzle_wrapper_collection_get_mapping(c *C.collection, options *C.query_options) *C.mapping_result {
	result := (*C.mapping_result)(C.calloc(1, C.sizeof_mapping_result))
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).GetMapping(opts)

	if err != nil {
		Set_mapping_result_error(result, err)
		return result
	}

	result.result = goToCMapping(res)

	return result
}
