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
  "unsafe"
  "time"
  "strconv"
  "github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.ack_response {
  result := (*C.ack_response)(C.calloc(1, C.sizeof_ack_response))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).CreateIndex(C.GoString(index), opts)

  if err != nil {
    Set_ack_response_error(result, err)
    return result
  }

  result.acknowledged = C.bool(res.Acknowledged)
  result.shards_acknowledged = C.bool(res.ShardsAcknowledged)

  return result
}

//export kuzzle_wrapper_get_auto_refresh
func kuzzle_wrapper_get_auto_refresh(k *C.kuzzle, index *C.char, options *C.query_options) *C.bool_result {
  result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetAutoRefresh(C.GoString(index), opts)
  if err != nil {
    Set_bool_result_error(result, err)
    return result
  }

  result.result = C.bool(res)

  return result
}

//export kuzzle_wrapper_refresh_index
func kuzzle_wrapper_refresh_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.shards {
  result := (*C.shards)(C.calloc(1, C.sizeof_shards))
  opts := SetQueryOptions(options)

  shards, err := (*kuzzle.Kuzzle)(k.instance).RefreshIndex(C.GoString(index), opts)
  if err != nil {
    Set_shards_error(result, err)
    return result
  }

  result.total = C.int(shards.Total)
  result.successful = C.int(shards.Successful)
  result.failed = C.int(shards.Failed)
  return result
}

//export kuzzle_wrapper_set_auto_refresh
func kuzzle_wrapper_set_auto_refresh(k *C.kuzzle, index *C.char, auto_refresh C.uint, options *C.query_options) *C.bool_result {
  result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
  opts := SetQueryOptions(options)

  autoRefresh := auto_refresh != 0

  res, err := (*kuzzle.Kuzzle)(k.instance).SetAutoRefresh(C.GoString(index), autoRefresh, opts)
  if err != nil {
    Set_bool_result_error(result, err)
    return result
  }

  result.result = C.bool(res)

  return result
}

//export kuzzle_wrapper_list_collections
func kuzzle_wrapper_list_collections(k *C.kuzzle, index *C.char, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).ListCollections(C.GoString(index), opts)
  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)

  buffer := C.CString(string(r))
  result.result = C.json_tokener_parse(buffer)
  C.free(unsafe.Pointer(buffer))

  return result
}

//export kuzzle_wrapper_list_indexes
func kuzzle_wrapper_list_indexes(k *C.kuzzle, options *C.query_options) *C.string_array_result {
  result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).ListIndexes(opts)
  if err != nil {
    Set_string_array_result_error(result, err)
    return result
  }

  result.result = (**C.char)(C.calloc(C.size_t(len(res)), C.sizeof_char_ptr))
  result.length = C.ulong(len(res))
  
  cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(result.result))[:len(res):len(res)]

  for i, substring := range res {
    cArray[i] = C.CString(substring)
  }

  return result
}

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(k *C.kuzzle, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)
  
  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)

  buffer := C.CString(string(r))
  defer C.free(unsafe.Pointer(buffer))

  result.result = C.json_tokener_parse(buffer)

  return result
}

//export kuzzle_wrapper_get_statistics
func kuzzle_wrapper_get_statistics(k *C.kuzzle, timestamp C.time_t, options *C.query_options) *C.statistics {
  result := (*C.statistics)(C.calloc(1, C.sizeof_statistics))
  opts := SetQueryOptions(options)

  t, _ := strconv.ParseInt(C.GoString(C.ctime(&timestamp)), 10, 64)
  tm := time.Unix(t, 0)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetStatistics(&tm, opts)

  if err != nil {
    Set_statistics_error(result, err)
    return result
  }

  ongoing, _ := json.Marshal(res.OngoingRequests)
  completed_requests, _ := json.Marshal(res.CompletedRequests)
  connections, _ := json.Marshal(res.Connections)
  failed_requests, _ := json.Marshal(res.FailedRequests)

  c_ongoing := C.CString(string(ongoing))
  c_completed_request := C.CString(string(completed_requests))
  c_connections := C.CString(string(connections))
  c_failed_requests := C.CString(string(failed_requests))

  result.ongoing_requests = C.json_tokener_parse(c_ongoing)
  result.completed_requests = C.json_tokener_parse(c_completed_request)
  result.connections = C.json_tokener_parse(c_connections)
  result.failed_requests = C.json_tokener_parse(c_failed_requests)
  result.timestamp = C.ulonglong(res.Timestamp)

  C.free(unsafe.Pointer(c_ongoing))
  C.free(unsafe.Pointer(c_completed_request))
  C.free(unsafe.Pointer(c_connections))
  C.free(unsafe.Pointer(c_failed_requests))

  return result
}

//export kuzzle_wrapper_get_server_info
func kuzzle_wrapper_get_server_info(k *C.kuzzle, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetServerInfo(opts)

  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

  r, _ := json.Marshal(res)
  buffer := C.CString(string(r))
  defer C.free(unsafe.Pointer(buffer))

  result.result = C.json_tokener_parse(buffer)

  return result
}

//export kuzzle_wrapper_now
func kuzzle_wrapper_now(k *C.kuzzle, options *C.query_options) *C.int_result {
  result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))
  opts := SetQueryOptions(options)

  time, err := (*kuzzle.Kuzzle)(k.instance).Now(opts)
  if err != nil {
    Set_int_result_error(result, err)
    return result
  }

  result.result = C.longlong(time)
  return result
}
