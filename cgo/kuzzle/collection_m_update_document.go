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

//export kuzzle_wrapper_collection_m_update_document
func kuzzle_wrapper_collection_m_update_document(c *C.collection, result *C.kuzzle_search_response, documents **C.document, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).MUpdateDocument(goDocuments(documents), opts)
	if err != nil {
		if err.Error() == "Collection.MUpdateDocument: please provide at least one document to update" {
			return C.int(C.EINVAL)
		}
		result.error = ToCString_2048(err.Error())
		return 0
	}

	go_to_c_search_result(res, result)

	return 0
}