package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_set_headers
func kuzzle_wrapper_set_headers(k *C.kuzzle, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		(*kuzzle.Kuzzle)(k.instance).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}
}
