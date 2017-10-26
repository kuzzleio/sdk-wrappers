package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_scroll_specifications
func kuzzle_wrapper_collection_scroll_specifications(c *C.collection, scrollId *C.char, options *C.query_options) *C.specification_search_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).ScrollSpecifications(C.GoString(scrollId), opts)

	return goToCSpecificationSearchResult(res, err)
}
