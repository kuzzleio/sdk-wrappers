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

// apply a types.KuzzleError on a ack_result* C struct
func Set_ack_result_error(s *C.ack_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a bool_result* C struct
func Set_bool_result_error(s *C.bool_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a kuzzle_response* C struct
func Set_kuzzle_response_error(s *C.kuzzle_response, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a statistics* C struct
func Set_statistics_error(s *C.statistics, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a string_array_result* C struct
func Set_string_array_result_error(s *C.string_array_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a login_result* C struct
func Set_login_result_error(s *C.login_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a int_result* C struct
func Set_int_result_error(s *C.int_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a string_result* C struct
func Set_string_result_error(s *C.string_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a shards* C struct
func Set_shards_result_error(s *C.shards_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a user* C struct
func Set_user_error(s *C.user, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a document* C struct
func Set_document_error(s *C.document_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}

// apply a types.KuzzleError on a kuzzle_search_result* C struct
func Set_kuzzle_search_result_error(s *C.kuzzle_search_result, err error) {
  kuzzleError := err.(*types.KuzzleError)
  s.status = C.int(kuzzleError.Status)
  s.error = C.CString(kuzzleError.Message)

  if len(kuzzleError.Stack) > 0 {
    s.stack = C.CString(kuzzleError.Stack)
  }
}
