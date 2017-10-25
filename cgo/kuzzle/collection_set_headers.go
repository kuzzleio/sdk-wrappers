package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_collection_set_headers
func kuzzle_wrapper_collection_set_headers(c *C.collection, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		cToGoCollection(c).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}
