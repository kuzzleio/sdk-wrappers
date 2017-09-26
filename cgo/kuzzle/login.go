package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
)

//export kuzzle_wrapper_login
func kuzzle_wrapper_login(k *C.Kuzzle, result *C.login_result, strategy *C.char, credentials *C.json_object, expires_in *C.int) C.int {
	jp := JsonParser{}
	jp.Parse(credentials)

	var expire int
	if expires_in != nil {
		expire = int(*expires_in)
	}
	res, err := (*kuzzle.Kuzzle)(k.instance).Login(C.GoString(strategy), jp.GetContent(), &expire)
	if err != nil {
		if err.Error() == "Kuzzle.Login: cannot authenticate to Kuzzle without an authentication strategy" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
	}

	result.jwt = *(*[512]C.char)(unsafe.Pointer(C.CString(res)))

	return 0
}
