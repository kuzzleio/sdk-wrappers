package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(k *C.Kuzzle, result *C.token_validity, token *C.char) C.int {
	res, err := (*kuzzle.Kuzzle)(k.instance).CheckToken(C.GoString(token))
	if err != nil {
		if err.Error() == "Kuzzle.CheckToken: token required" {
			return C.int(C.EINVAL)
		} else {
			result.error = ToCString_2048(err.Error())
			return 0
		}
	}

	var valid C.uint

	if res.Valid {
		valid = 1
	} else {
		valid = 0
	}

	*result = C.token_validity{
		valid:     valid,
		state:     ToCString_512(res.State),
		expiresAt: C.int(res.ExpiresAt),
	}

	return 0
}
