package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_set_headers
<<<<<<< HEAD
func kuzzle_wrapper_set_headers(k *C.Kuzzle, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		(*kuzzle.Kuzzle)(k.instance).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}
=======
func kuzzle_wrapper_set_headers(k *C.kuzzle, content *C.json_object, replace C.uint) {
	jp := JsonParser{}
	jp.Parse(content)

	r := replace != 0
	(*kuzzle.Kuzzle)(k.instance).SetHeaders(jp.GetContent(), r)
>>>>>>> origin/master
}
