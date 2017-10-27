package main

/*
  #cgo CFLAGS: -I../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include <string.h>
  #include <json-c/json.h>
  #include "kuzzle.h"
  #include "sdk_wrappers_internal.h"
*/
import "C"
import (
  "unsafe"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_ms_append
func kuzzle_wrapper_ms_append(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Append(
    C.GoString(key), 
    C.GoString(value), 
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_bitcount
func kuzzle_wrapper_ms_bitcount(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitcount(
    C.GoString(key), 
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_bitop
func kuzzle_wrapper_ms_bitop(k *C.kuzzle, key *C.char, operation *C.char, keys **C.char, klen C.uint, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitop(
    C.GoString(key), 
    C.GoString(operation),
    cToGoStrings(keys, klen),
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_bitpos
func kuzzle_wrapper_ms_bitpos(k *C.kuzzle, key *C.char, bit C.uchar, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitpos(
    C.GoString(key), 
    int(bit),
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_dbsize
func kuzzle_wrapper_ms_dbsize(k *C.kuzzle, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Dbsize(SetQueryOptions(options))
  
  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_decr
func kuzzle_wrapper_ms_decr(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Decr(
    C.GoString(key), 
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_decrby
func kuzzle_wrapper_ms_decrby(k *C.kuzzle, key *C.char, value C.int, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Decrby(
    C.GoString(key), 
    int(value), 
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_del
func kuzzle_wrapper_ms_del(k *C.kuzzle, keys **C.char, klen C.uint, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Del(
    cToGoStrings(keys, klen),
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_exists
func kuzzle_wrapper_ms_exists(k *C.kuzzle, keys **C.char, klen C.uint, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Exists(
    cToGoStrings(keys, klen),
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_expire
func kuzzle_wrapper_ms_expire(k *C.kuzzle, key *C.char, seconds C.ulong, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Expire(
    C.GoString(key),
    int(seconds),
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_expireat
func kuzzle_wrapper_ms_expireat(k *C.kuzzle, key *C.char, ts C.ulonglong, options *C.query_options) *C.int_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Expireat(
    C.GoString(key),
    int(ts),
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_flushdb
func kuzzle_wrapper_ms_flushdb(k *C.kuzzle, options *C.query_options) *C.string_result {
  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Flushdb(SetQueryOptions(options))

  return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_geoadd
func kuzzle_wrapper_ms_geoadd(k *C.kuzzle, key *C.char, points **C.json_object, plen C.uint, options *C.query_options) *C.int_result {
  wrapped := (*[1 << 30]*C.json_object)(unsafe.Pointer(points))[:plen:plen]
  gopoints := make([]*types.GeoPoint, int(plen))

  for i, jobj := range(wrapped) {
    stringified := C.json_object_to_json_string(jobj)
    gobytes := C.GoBytes(unsafe.Pointer(stringified), C.int(C.strlen(stringified)))
    json.Unmarshal(gobytes, gopoints[i])
  }

  res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Geoadd(
    C.GoString(key),
    gopoints,
    SetQueryOptions(options))

  return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_geodist
func kuzzle_wrapper_ms_geodist(k *C.kuzzle, key *C.char, member1 *C.char, member2 *C.char, options C.query_options*) {

}
