package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(k *C.Kuzzle, token *C.char) *C.token_validity {
	result := (*C.token_validity)(C.calloc(1, C.sizeof_token_validity))

	res, err := (*kuzzle.Kuzzle)(k.instance).CheckToken(C.GoString(token))
	if err != nil {
		Set_token_validity_error(result, err)
		return result
	}

	if res.Valid {
		result.valid = 1
	} else {
		result.valid = 0
	}

	result.status = 200
	result.state = C.CString(res.State)
	result.expiresAt = C.longlong(res.ExpiresAt)

	return result
}
