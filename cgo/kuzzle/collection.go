package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"unsafe"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_new_collection
func kuzzle_wrapper_new_collection(k *C.Kuzzle, colName *C.char, index *C.char) *C.collection {
	// TODO Must be freed in C
	col := (*C.collection)(C.calloc(1, C.sizeof_collection))
	// TODO Must be freed in C
	C.strcpy(col.index, index)
	// TODO Must be freed in C
	C.strcpy(col.collection, colName)

	// TODO Must be freed by the Kuzzle destructor
	col.kuzzle = unsafe.Pointer(k)

	return col
}

// TODO check if it is still legit - refactor
func goToCSearchResult(goRes *collection.SearchResult, cRes *C.kuzzle_search_result) {
	cRes.result.total = C.int(goRes.Total)

	if len(goRes.Hits) > 0 {
		// TODO Refactor
		hits := make([]*C.document, len(goRes.Hits) + 1)

		for i := 0; i < len(goRes.Hits); i++ {
			var doc C.document
			// TODO register it in global
			t := goRes.Hits[i]
			doc.instance = unsafe.Pointer(&t)
			hits[i] = &doc
		}
		hits[len(goRes.Hits)] = nil

		cRes.result.hits = &hits[0]
	}
}

// TODO check if it is still legit - refactor
func goToCSpecificationSearchResult(goRes *types.KuzzleSpecificationSearchResult, cRes *C.kuzzle_specification_search_result) {
	cRes.result.total = C.int(goRes.Total)

	if len(goRes.Hits) > 0 {
		hits := make([]*C.kuzzle_specification, len(goRes.Hits) + 1)

		for i := 0; i < len(goRes.Hits); i++ {
			var spec C.kuzzle_specification
			// TODO register it in global
			t := goRes.Hits[i]
			spec.instance = unsafe.Pointer(&t)
			hits[i] = &spec
		}
		hits[len(goRes.Hits)] = nil

		cRes.result.hits = &hits[0]
	}
}

// TODO check if it is still legit - refactor
func cToGoSearchFilters(searchFilters *C.search_filters) *types.SearchFilters {
	return &types.SearchFilters{
		Query: JsonCConvert(searchFilters.query),
		Sort: JsonCConvert(searchFilters.sort).([]interface{}),
		Aggregations: JsonCConvert(searchFilters.aggregations),
		SearchAfter: JsonCConvert(searchFilters.search_after).([]interface{}),
	}
}