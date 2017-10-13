package main

/*
  #cgo CFLAGS: -I../../../headers
  #cgo LDFLAGS: -ljson-c

  #include "kuzzle.h"
*/
import "C"

import (
  "github.com/kuzzleio/sdk-go/types"
)

// apply a types.KuzzleError on a json_result* C struct
func Set_json_result_error(s *C.json_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a token_validity* C struct
func Set_token_validity_error(s *C.token_validity, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a ack_response* C struct
func Set_ack_response_error(s *C.ack_response, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a ack_response* C struct
func Set_bool_result_error(s *C.bool_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}
