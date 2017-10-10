package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_collection_scroll_specifications
func kuzzle_wrapper_collection_scroll_specifications(c *C.collection, result *C.kuzzle_specification_search_response, scrollId *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).ScrollSpecifications(C.GoString(scrollId), opts)
	if err != nil {
		if err.Error() == "Collection.ScrollSpecifications: scroll id required" {
			return C.int(C.EINVAL)
		}
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return 0
	}

	go_to_c_specification_search_result(&res, result)

	return 0
}