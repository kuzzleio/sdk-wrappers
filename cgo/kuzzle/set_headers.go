package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_set_headers
func kuzzle_wrapper_set_headers(k *C.Kuzzle, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		var r bool

		if replace == 1 {
			r = true
		}

		(*kuzzle.Kuzzle)(k.instance).SetHeaders((types.HeadersData)(JsonCConvert(content).(map[string]interface{})), r)
	}

	return
}
