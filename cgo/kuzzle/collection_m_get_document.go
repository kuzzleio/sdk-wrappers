package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_collection_m_get_document
func kuzzle_wrapper_collection_m_get_document(c *C.collection, ids **C.char, idsCount C.uint, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)
	gIds := cToGoStrings(ids, idsCount)
	res, err := cToGoCollection(c).MGetDocument(gIds, opts)

	return goToCSearchResult(c, res, err)
}