package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <json/json.h>
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

//export kuzzle_wrapper_create_my_credentials
func kuzzle_wrapper_create_my_credentials(result *C.json_object, strategy *C.char, credentials *C.json_object, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := KuzzleInstance.CreateMyCredentials(C.GoString(strategy), credentials, opts)
	if err != nil && err.Error() == "Kuzzle.CreateMyCredentials: strategy is required" {
		return C.int(C.EINVAL)
	}

	var jsonRes *C.json_object
	r, _ := json.Marshal(res)

	jsonRes = C.json_tokener_parse(C.CString(string(r)))
	result = jsonRes

	return 0
}
