package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include <kuzzle.h>
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(k *C.Kuzzle, token *C.char) *C.token_validity {
	res, err := (*kuzzle.Kuzzle)(k.instance).CheckToken(C.GoString(token))
	result := (*C.token_validity)(C.calloc(1, C.sizeof_token_validity))

	if err != nil {
			result.error = C.CString(err.Error())
			return result
	}

	if res.Valid {
		result.valid = 1
	} else {
		result.valid = 0
	}

	result.state = C.CString(res.State)
	result.expiresAt = C.int(res.ExpiresAt)

	return result
}
