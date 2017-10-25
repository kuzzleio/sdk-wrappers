package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_collection_m_update_document
func kuzzle_wrapper_collection_m_update_document(c *C.collection, documents **C.document, docCount C.int, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)
	res, err := cToGoCollection(c).MUpdateDocument(cToGoDocuments(c, documents, docCount), opts)

	return goToCSearchResult(c, res, err)
}
