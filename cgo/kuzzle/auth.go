package main

/*
  #cgo CFLAGS: -I../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include "kuzzle.h"
  #include "sdk_wrappers_internal.h"
*/
import "C"
import (
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "unsafe"
  "github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_set_jwt
func kuzzle_wrapper_set_jwt(k *C.kuzzle, token *C.char) {
  (*kuzzle.Kuzzle)(k.instance).SetJwt(C.GoString(token))
}

//export kuzzle_wrapper_unset_jwt
func kuzzle_wrapper_unset_jwt(k *C.kuzzle) {
  (*kuzzle.Kuzzle)(k.instance).UnsetJwt()
}

// Allocates memory
//export kuzzle_wrapper_get_jwt
func kuzzle_wrapper_get_jwt(k *C.kuzzle) *C.char {
  return C.CString((*kuzzle.Kuzzle)(k.instance).GetJwt())
}

//export kuzzle_wrapper_login
func kuzzle_wrapper_login(k *C.kuzzle, strategy *C.char, credentials *C.json_object, expires_in *C.int) *C.login_result {
  result := (*C.login_result)(C.calloc(1, C.sizeof_login_result))

  var expire int
  if expires_in != nil {
    expire = int(*expires_in)
  }

  res, err := (*kuzzle.Kuzzle)(k.instance).Login(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), &expire)

  if err != nil {
    Set_login_result_error(result, err)
    return result
  }

  result.jwt = C.CString(res)

  return result
}

//export kuzzle_wrapper_logout
func kuzzle_wrapper_logout(k *C.kuzzle) *C.char {
  err := (*kuzzle.Kuzzle)(k.instance).Logout()
  if err != nil {
    return C.CString(err.Error())
  }

  return nil
}


//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(k *C.kuzzle, token *C.char) *C.token_validity {
  result := (*C.token_validity)(C.calloc(1, C.sizeof_token_validity))

  res, err := (*kuzzle.Kuzzle)(k.instance).CheckToken(C.GoString(token))
  if err != nil {
    Set_token_validity_error(result, err)
    return result
  }

  result.valid = C.bool(res.Valid)

  result.status = 200
  result.state = C.CString(res.State)
  result.expires_at = C.ulonglong(res.ExpiresAt)

  return result
}

//export kuzzle_wrapper_create_my_credentials
func kuzzle_wrapper_create_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  var opts types.QueryOptions
  opts = SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).CreateMyCredentials(C.GoString(strategy),JsonCConvert(credentials).(map[string]interface{}), opts)
  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)

  buffer := C.CString(string(r))
  defer C.free(unsafe.Pointer(buffer))

  result.result.ptr = C.json_tokener_parse(buffer)

  return result
}


//export kuzzle_wrapper_delete_my_credentials
func kuzzle_wrapper_delete_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.ack_result {
  result := (*C.ack_result)(C.calloc(1, C.sizeof_ack_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).DeleteMyCredentials(C.GoString(strategy), opts)

  if err != nil {
    Set_ack_result_error(result, err)
    return result
  }

  result.acknowledged = C.bool(res.Acknowledged)
  result.shards_acknowledged = C.bool(res.ShardsAcknowledged)

  return result
}

//export kuzzle_wrapper_get_my_credentials
func kuzzle_wrapper_get_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetMyCredentials(C.GoString(strategy), opts)

  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)
  buffer := C.CString(string(r))
  defer C.free(unsafe.Pointer(buffer))

  result.result.ptr = C.json_tokener_parse(buffer)

  return result
}

//export kuzzle_wrapper_update_my_credentials
func kuzzle_wrapper_update_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).UpdateMyCredentials(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), opts)
  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)
  buffer := C.CString(string(r))
  result.result.ptr = C.json_tokener_parse(buffer)
  C.free(unsafe.Pointer(buffer))

  return result
}

//export kuzzle_wrapper_validate_my_credentials
func kuzzle_wrapper_validate_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.bool_result {
  result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).ValidateMyCredentials(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), opts)
  if err != nil {
    Set_bool_result_error(result, err)
    return result
  }

  result.result = C.bool(res)

  return result
}

//export kuzzle_wrapper_get_my_rights
func kuzzle_wrapper_get_my_rights(k *C.kuzzle, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetMyRights(opts)

  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)
  buffer := C.CString(string(r))
  result.result.ptr = C.json_tokener_parse(buffer)

  C.free(unsafe.Pointer(buffer))
  return result
}

//export kuzzle_wrapper_update_self
func kuzzle_wrapper_update_self(k *C.kuzzle, credentials *C.json_object, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).UpdateSelf(JsonCConvert(credentials).(map[string]interface{}), opts)
  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)
  buffer := C.CString(string(r))
  result.result.ptr = C.json_tokener_parse(buffer)
  C.free(unsafe.Pointer(buffer))

  return result
}

//export kuzzle_wrapper_who_am_i
func kuzzle_wrapper_who_am_i(k *C.kuzzle) *C.user {
  user := (*C.user)(C.calloc(1, C.sizeof_user))

  res, err := (*kuzzle.Kuzzle)(k.instance).WhoAmI()
  if err != nil {
    Set_user_error(user, err)
    return user
  }

  user.meta = goToCKuzzleMeta(res.Meta)

  buffer := C.CString(string(res.Source))
  user.source = C.json_tokener_parse(buffer)
  C.free(unsafe.Pointer(buffer))

  user.strategies_length = C.ulong(len(res.Strategies))
  user.strategies = (**C.char)(C.calloc(C.size_t(user.strategies_length), C.sizeof_char_ptr))
  cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(user.strategies))[:len(res.Strategies):len(res.Strategies)]

  for i, substring := range res.Strategies {
    cArray[i] = C.CString(substring)
  }

  user.id = C.CString(res.Id)

  return user
}
