package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
  "unsafe"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/security"
)

//export kuzzle_wrapper_security_fetch_profile
func kuzzle_wrapper_security_fetch_profile(s *C.security, result *C.kuzzle_response, id *C.char, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetQueryOptions(options)
  }

  res, err := (*security.Security)(s.instance).Profile.Fetch(C.GoString(id), opts)
  if err != nil {
    if err.Error() == "Security.Profile.Fetch: profile id required" {
      return C.int(C.EINVAL)
    } else {
      result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
      return 0
    }
  }

  var jsonRes *C.json_object
  r, _ := json.Marshal(res)

  jsonRes = C.json_tokener_parse(C.CString(string(r)))
  result.result = jsonRes

  return 0
}