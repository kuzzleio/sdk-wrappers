package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_set_headers
func kuzzle_wrapper_set_headers(k *C.Kuzzle, content *C.json_object, replace C.uint) {
	jp := JsonParser{}
	jp.Parse(content)

	var r bool
	if replace == 1 {
		r = true
	}
	(*kuzzle.Kuzzle)(k.instance).SetHeaders(jp.GetContent(), r)
}
