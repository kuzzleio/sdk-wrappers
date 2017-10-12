package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_login
func kuzzle_wrapper_login(k *C.Kuzzle, result *C.login_result, strategy *C.char, credentials *C.json_object, expires_in *C.int) C.int {

	var expire int
	if expires_in != nil {
		expire = int(*expires_in)
	}
	res, err := (*kuzzle.Kuzzle)(k.instance).Login(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), &expire)
	if err != nil {
		if err.Error() == "Kuzzle.Login: cannot authenticate to Kuzzle without an authentication strategy" {
			return C.int(C.EINVAL)
		} else {
			result.error = ToCString_2048(err.Error())
			return 0
		}
	}

	var arr [512]C.char

	for i, v := range res {
		arr[i] = (C.char)(v)
	}

	result.jwt = arr

	return 0
}
