package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_m_replace_document
func kuzzle_wrapper_collection_m_replace_document(c *C.collection, result *C.kuzzle_search_result, documents **C.document, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MReplaceDocument(goDocuments(documents), opts)

	if err != nil {
		if err.Error() == "Collection.MReplaceDocument: please provide at least one document to replace" {
			return C.int(C.EINVAL)
		}
		result.error = ToCString_2048(err.Error())
		return 0
	}

	goToCSearchResult(res, result)

	return 0
}
