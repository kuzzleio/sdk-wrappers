package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import "github.com/kuzzleio/sdk-go/collection"

//export kuzzle_wrapper_collection_set_headers
func kuzzle_wrapper_collection_set_headers(c *C.collection, content *C.json_object, replace C.uint) {
	jp := JsonParser{}
	jp.Parse(content)

	var r bool
	if replace == 1 {
		r = true
	}

	(*collection.Collection)(c.instance).SetHeaders(jp.GetContent(), r)

	return
}
