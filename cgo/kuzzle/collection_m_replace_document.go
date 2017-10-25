package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_m_replace_document
func kuzzle_wrapper_collection_m_replace_document(c *C.collection, documents **C.document, docCount C.int, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)
	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MReplaceDocument(cToGoDocuments(documents, docCount, c), opts)

	return goToCSearchResult(res, c, err)
}
