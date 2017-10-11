package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"fmt"
)

//export kuzzle_wrapper_collection_count
func kuzzle_wrapper_collection_count(c *C.collection, result *C.int_response, searchFilters *C.search_filters, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	// TODO: Initialize Search Filters from C to Go
	filters := types.SearchFilters{}

	res, err := (*collection.Collection)(c.instance).Count(&filters, opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	result.result = C.int(res)
}

func go_to_c_search_filter (cSf *C.search_filters) *types.SearchFilters {
	filters := types.SearchFilters{}
	jp := JsonParser{}

	if cSf.query != nil {
		jp.Parse(cSf.query)
		filters.Query = jp.GetContent()
	}

	if cSf.sort != nil {
		jp.Parse(cSf.sort)
		//filters.Sort = jp.GetContent()
	}

	return &filters
}

//export test_me
func test_me (filter *C.json_object) {
	jp := JsonParser{}
	jp.Parse(filter)

	fmt.Println(jp.GetContent())
}