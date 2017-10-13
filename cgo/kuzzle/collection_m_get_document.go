package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_collection_m_get_document
func kuzzle_wrapper_collection_m_get_document(c *C.collection, result *C.kuzzle_search_response, ids **C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).MGetDocument(goStrings(ids), opts)
	if err != nil {
		if err.Error() == "Collection.MGetDocument: please provide at least one id of document to retrieve" {
			return C.int(C.EINVAL)
		}
		result.error = ToCString_2048(err.Error())
		return 0
	}

	goToCSearchResult(res, result)

	return 0
}