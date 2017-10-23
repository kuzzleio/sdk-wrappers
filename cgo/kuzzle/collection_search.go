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

//export kuzzle_wrapper_collection_search
// TODO
func kuzzle_wrapper_collection_search(c *C.collection, result *C.kuzzle_search_result, searchFilters *C.search_filters, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.Search(cToGoSearchFilters(searchFilters), opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	goToCSearchResult(res, result)
}