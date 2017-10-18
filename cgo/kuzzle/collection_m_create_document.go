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

//export kuzzle_wrapper_collection_m_create_document
func kuzzle_wrapper_collection_m_create_document(c *C.collection, result *C.kuzzle_search_response, documents **C.document, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))

	res, err := col.MCreateDocument(goDocuments(documents), opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	goToCSearchResult(res, result)
}

//export kuzzle_wrapper_collection_m_create_or_replace_document
func kuzzle_wrapper_collection_m_create_or_replace_document(c *C.collection, result *C.kuzzle_search_response, documents **C.document, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MCreateOrReplaceDocument(goDocuments(documents), opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	goToCSearchResult(res, result)
}
