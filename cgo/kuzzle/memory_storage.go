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
  "unsafe"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_ms_append
func kuzzle_wrapper_ms_append(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Append(C.GoString(key), C.GoString(value), opts)

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_bitcount
func kuzzle_wrapper_ms_bitcount(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
  opts := SetQueryOptions(options)
  
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitcount(C.GoString(key), opts)

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_bitop
func kuzzle_wrapper_ms_bitop(k *C.kuzzle, key *C.char, operation *C.char, keys **C.char, klen C.int, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitop(
    C.GoString(key), 
    C.GoString(operation),
    toStringSlice(keys, klen),
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_bitpos
func kuzzle_wrapper_ms_bitpos(k *C.kuzzle, key *C.char, bit C.uchar, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitpos(
    C.GoString(key), 
    int(bit),
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_dbsize
func kuzzle_wrapper_ms_dbsize(k *C.kuzzle, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Dbsize(SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_decr
func kuzzle_wrapper_ms_decr(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
  opts := SetQueryOptions(options)
  
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Decr(C.GoString(key), opts)

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_decrby
func kuzzle_wrapper_ms_decrby(k *C.kuzzle, key *C.char, value C.int, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
  
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Decrby(
    C.GoString(key), 
    int(value), 
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_del
func kuzzle_wrapper_ms_del(k *C.kuzzle, keys **C.char, klen C.int, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Del(
    toStringSlice(keys, klen),
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_exists
func kuzzle_wrapper_ms_exists(k *C.kuzzle, keys **C.char, klen C.int, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Exists(
    toStringSlice(keys, klen),
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_expire
func kuzzle_wrapper_ms_expire(k *C.kuzzle, key *C.char, seconds C.ulong, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Expire(
    C.GoString(key),
    int(seconds),
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_expireat
func kuzzle_wrapper_ms_expireat(k *C.kuzzle, key *C.char, ts C.ulonglong, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Expireat(
    C.GoString(key),
    int(ts),
    SetQueryOptions(options))

  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(res)
  return result
}

//export kuzzle_wrapper_ms_flushdb
func kuzzle_wrapper_ms_flushdb(k *C.kuzzle, options *C.query_options) *C.string_result {
  result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Flushdb(SetQueryOptions(options))

  if err != nil {
    Set_string_result_error(result, err)
    return result
  }

  result.result = C.CString(res)
  return result
}

//export kuzzle_wrapper_ms_geoadd
func kuzzle_wrapper_ms_geoadd(k *C.kuzzle, key *C.char, points **C.json_object, int plen, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

  wrapped := (*[1 << 30]*C.json_object)(unsafe.Pointer(points))[:plen:plen]
  slice := make([]*types.GeoPoint, plen)

  for i := 0; i < int(plen); i++ {
    json_object_object_get_ex
    slice[i].Lon = wrapped[i].
  }
}

func toStringSlice(arr **C.char, arr_len C.int) []string {
  wrapped := (*[1 << 30]C.char_ptr)(unsafe.Pointer(arr))[:arr_len:arr_len]
  slice := make([]string, arr_len)

  for i := 0; i < int(arr_len); i++ {
    slice[i] = C.GoString(wrapped[i])
  }

  return slice
}
