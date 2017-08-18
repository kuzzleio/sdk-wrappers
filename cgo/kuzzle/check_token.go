package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
 */
import "C"
import "unsafe"

//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(result *C.token_validity, token *C.char) C.int {
	res, err := KuzzleInstance.CheckToken(C.GoString(token))
	if err != nil && err.Error() == "Kuzzle.CheckToken: token required" {
		return C.int(C.EINVAL)
	}

	var valid C.uint

	if res.Valid {
		valid = 1
	} else {
		valid = 0
	}

	*result = C.token_validity{
		valid:     valid,
		state: *(*[512]C.char)(unsafe.Pointer(C.CString(res.State))),
		expiresAt: C.int(res.ExpiresAt),
	}

	return 0
}
