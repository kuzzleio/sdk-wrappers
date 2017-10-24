package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import "github.com/kuzzleio/sdk-go/kuzzle"

//export kuzzle_wrapper_set_jwt
func kuzzle_wrapper_set_jwt(k *C.kuzzle, token *C.char) {
	(*kuzzle.Kuzzle)(k.instance).SetJwt(C.GoString(token))
}
