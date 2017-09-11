package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_set_headers
func kuzzle_wrapper_set_headers(content *C.json_object, replace C.uint) {
	jp := JsonParser{}
	jp.Parse(content)

	var r bool
	if replace == 1 {
		r = true
	}
	KuzzleInstance.SetHeaders(jp.GetContent(), r)
}