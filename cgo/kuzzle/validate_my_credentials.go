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

//export kuzzle_wrapper_validate_my_credentials
func kuzzle_wrapper_validate_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
	opts := SetQueryOptions(options)

	jp := JsonParser{}
	jp.Parse(credentials)

	res, err := (*kuzzle.Kuzzle)(k.instance).ValidateMyCredentials(C.GoString(strategy), jp.GetContent(), opts)
	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(res)

	return result
}
