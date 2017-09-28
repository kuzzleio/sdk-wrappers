package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"fmt"
)

//export kuzzle_wrapper_create_my_credentials
func kuzzle_wrapper_create_my_credentials(k *C.Kuzzle, result *C.json_result, strategy *C.char, credentials *C.json_object, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	jp := JsonParser{}
	jp.Parse(credentials)

	res, err := (*kuzzle.Kuzzle)(k.instance).CreateMyCredentials(C.GoString(strategy), jp.GetContent(), opts)
	if err != nil {
		if err.Error() == "Kuzzle.CreateMyCredentials: strategy is required" {
			return C.int(C.EINVAL)
		} else {
			result.error = ToCString_2048(err.Error())
			fmt.Printf("%s\n", err.Error())
			return 0
		}
	}

	var jsonRes *C.json_object
	r, _ := json.Marshal(res)

	jsonRes = C.json_tokener_parse(C.CString(string(r)))
	result.result = jsonRes

	return 0
}
