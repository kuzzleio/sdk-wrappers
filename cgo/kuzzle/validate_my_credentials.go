package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_validate_my_credentials
func kuzzle_wrapper_validate_my_credentials(k *C.Kuzzle, result *C.bool_result, strategy *C.char, credentials *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	jp := JsonParser{}
	jp.Parse(credentials)

	res, err := (*kuzzle.Kuzzle)(k.instance).ValidateMyCredentials(C.GoString(strategy), jp.GetContent(), opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
	}

	var r C.uint
	if res {
		r = 1
	} else {
		r = 0
	}
	result.result = r
}
