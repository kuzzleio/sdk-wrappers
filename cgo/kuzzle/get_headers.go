package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import "encoding/json"

//export kuzzle_wrapper_get_headers
func kuzzle_wrapper_get_headers() *C.json_object {
	res := KuzzleInstance.GetHeaders()
	r, _ := json.Marshal(res)

	return C.json_tokener_parse(C.CString(string(r)))
}