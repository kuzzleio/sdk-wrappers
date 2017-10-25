package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_collection_m_create_document
func kuzzle_wrapper_collection_m_create_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.search_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).MCreateDocument(cToGoDocuments(c, documents, docCount), opts)

	 return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_m_create_or_replace_document
func kuzzle_wrapper_collection_m_create_or_replace_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.search_result {
	opts := SetQueryOptions(options)

	col := cToGoCollection(c)
	res, err := col.MCreateOrReplaceDocument(cToGoDocuments(c, documents, docCount), opts)

	return goToCSearchResult(c, res, err)
}
