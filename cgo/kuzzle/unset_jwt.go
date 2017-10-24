package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_unset_jwt
func kuzzle_wrapper_unset_jwt(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).UnsetJwt()
}
