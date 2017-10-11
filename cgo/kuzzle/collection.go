package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_new_collection
func kuzzle_wrapper_new_collection(c *C.collection, k *C.Kuzzle, colName *C.char, index *C.char) *C.collection {
	instance := collection.NewCollection((*kuzzle.Kuzzle)(k.instance), C.GoString(colName), C.GoString(index))

	c.instance = unsafe.Pointer(instance)

	return c
}

func go_to_c_search_result(goRes *collection.SearchResult, cRes *C.kuzzle_search_response) {
	cRes.result.total = C.int(goRes.Total)

	if len(goRes.Hits) > 0 {
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

func go_to_c_specification_search_result(goRes *types.KuzzleSpecificationSearchResult, cRes *C.kuzzle_specification_search_response) {
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