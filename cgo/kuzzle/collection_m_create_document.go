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

//export kuzzle_wrapper_collection_m_create_document
func kuzzle_wrapper_collection_m_create_document(c *C.collection, documents **C.document, docCount C.int, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))

	res, err := col.MCreateDocument(cToGoDocuments(documents, docCount, c), opts)

	 return goToCSearchResult(res, c, err)
}

//export kuzzle_wrapper_collection_m_create_or_replace_document
func kuzzle_wrapper_collection_m_create_or_replace_document(c *C.collection, documents **C.document, docCount C.int, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MCreateOrReplaceDocument(cToGoDocuments(documents, docCount, c), opts)

	return goToCSearchResult(res, c, err)
}
