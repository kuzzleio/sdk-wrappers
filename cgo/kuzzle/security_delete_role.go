package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
  "unsafe"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/security"
)

//export kuzzle_wrapper_security_delete_role
func kuzzle_wrapper_security_delete_role(s *C.security, result *C.kuzzle_string_response, id *C.char, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetQueryOptions(options)
  }

  res, err := (*security.Security)(s.instance).Role.Delete(C.GoString(id), opts)
  if err != nil {
    if err.Error() == "Security.Role.Delete: role id required" {
      return C.int(C.EINVAL)
    } else {
      result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
      return 0
    }
  }

  result.result = *(*[512]C.char)(unsafe.Pointer(C.CString(res)))

  return 0
}