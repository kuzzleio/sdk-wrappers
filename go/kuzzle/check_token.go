package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
 */
import "C"

//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(token *C.char) C.token_validity {
	res, err := KuzzleInstance.CheckToken(C.GoString(token))
	if err != nil {
		return C.token_validity{}
	}

	var valid C.char

	if res.Valid {
		valid = '1'
	} else {
		valid = '0'
	}

	return C.token_validity{
		valid: valid,
		state: C.CString(res.State),
		expiresAt: C.int(res.ExpiresAt),
	}
}
