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

//export kuzzle_wrapper_collection_m_get_document
func kuzzle_wrapper_collection_m_get_document(c *C.collection, ids **C.char, idsCount C.uint, options *C.query_options) *C.kuzzle_search_result {
	opts := SetQueryOptions(options)
	gIds := cToGoStrings(ids, idsCount)
	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MGetDocument(gIds, opts)

	return goToCSearchResult(res, c, err)
}