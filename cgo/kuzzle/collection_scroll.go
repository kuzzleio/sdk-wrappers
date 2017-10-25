package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"

//export kuzzle_wrapper_collection_scroll
func kuzzle_wrapper_collection_scroll(c *C.collection, scrollId *C.char, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).Scroll(C.GoString(scrollId), opts)

	return goToCSearchResult(c, res, err)
}
