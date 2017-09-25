package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_collection_search
func kuzzle_wrapper_collection_search(c *C.collection, result *C.kuzzle_search_response, search_filters *C.search_filters, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	// TODO: Initialize Search Filters from C to Go
	filters := types.SearchFilters{}

	res, err := (*collection.Collection)(c.instance).Search(filters, opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	result.result.total = C.int(res.Total)

	if len(res.Hits) > 0 {
		var hits *[len(res.Hits)]C.document

		for i := 0; i < len(res.Hits); i++ {
			var doc *C.document
			*doc.instance = unsafe.Pointer(res.Hits[i])
			hits[i] = doc
		}

		result.result.hits = hits
	}
}
