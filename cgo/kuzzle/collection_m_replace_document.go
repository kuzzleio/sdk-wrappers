package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_m_replace_document
func kuzzle_wrapper_collection_m_replace_document(c *C.collection, documents **C.document, options *C.query_options) *C.kuzzle_search_result {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MReplaceDocument(cToGoDocuments(documents, c), opts)

	return goToCSearchResult(res, c, err)
}
