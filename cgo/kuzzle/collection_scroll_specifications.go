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
	"github.com/kuzzleio/sdk-go/kuzzle"
)
//export kuzzle_wrapper_collection_scroll_specifications
// TODO
func kuzzle_wrapper_collection_scroll_specifications(c *C.collection, result *C.specification_search_result, scrollId *C.char, options *C.query_options) C.int {
	opts := SetQueryOptions(options)
	col := cToGoCollection(c)
	res, err := col.ScrollSpecifications(C.GoString(scrollId), opts)

	if err != nil {
		if err.Error() == "Collection.ScrollSpecifications: scroll id required" {
			return C.int(C.EINVAL)
		}
		result.error = ToCString_2048(err.Error())
		return 0
	}

	goToCSpecificationSearchResult(res, result)

	return 0
}
*/