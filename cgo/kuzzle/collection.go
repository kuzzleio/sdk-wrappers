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
)

//export kuzzle_wrapper_new_collection
func kuzzle_wrapper_new_collection(c *C.collection, k *C.Kuzzle, colName *C.char, index *C.char) *C.collection {
	instance := collection.NewCollection((*kuzzle.Kuzzle)(k.instance), C.GoString(colName), C.GoString(index))

	c.instance = unsafe.Pointer(instance)

	return c
}

func go_to_c_search_result(goRes *C.void, cRes *C.kuzzle_search_response) {
	res := (*collection.SearchResult)(goRes)

	cRes.result.total = C.int(res.Total)

	if len(res.Hits) > 0 {
		var hits *[len(res.Hits)]C.document

		for i := 0; i < len(res.Hits); i++ {
			var doc *C.document
			*doc.instance = unsafe.Pointer(res.Hits[i])
			hits[i] = doc
		}

		cRes.result.hits = hits
	}
}
