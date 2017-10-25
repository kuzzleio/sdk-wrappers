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

//export kuzzle_wrapper_collection_set_headers
func kuzzle_wrapper_collection_set_headers(c *C.collection, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
		col.SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}
