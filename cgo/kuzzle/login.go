package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_login
func kuzzle_wrapper_login(k *C.kuzzle, strategy *C.char, credentials *C.json_object, expires_in *C.int) *C.login_result {
	result := (*C.login_result)(C.calloc(1, C.sizeof_login_result))

	jp := JsonParser{}
	jp.Parse(credentials)

	var expire int
	if expires_in != nil {
		expire = int(*expires_in)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).Login(C.GoString(strategy), jp.GetContent(), &expire)
	if err != nil {
		Set_login_result_error(result, err)
		return result
	}

	result.jwt = C.CString(res)

	return result
}
