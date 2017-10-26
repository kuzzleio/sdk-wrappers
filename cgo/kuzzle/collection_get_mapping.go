package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_get_mapping
func kuzzle_wrapper_collection_get_mapping(c *C.collection, options *C.query_options) *C.mapping_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).GetMapping(opts)

	return goToCMappingResult(c, res, err)
}
