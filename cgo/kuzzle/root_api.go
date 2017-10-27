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
  "strconv"
  "time"
  "unsafe"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_list_collections
func kuzzle_wrapper_list_collections(k *C.kuzzle, index *C.char, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).ListCollections(C.GoString(index), opts)
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

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.ack_result {
  result := (*C.ack_result)(C.calloc(1, C.sizeof_ack_result))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).CreateIndex(C.GoString(index), opts)

  if err != nil {
    Set_ack_result_error(result, err)
    return result
  }

  result.acknowledged = C.bool(res.Acknowledged)
  result.shards_acknowledged = C.bool(res.ShardsAcknowledged)

  return result
}

//export kuzzle_wrapper_refresh_index
func kuzzle_wrapper_refresh_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.shards_result {
  result := (*C.shards_result)(C.calloc(1, C.sizeof_shards_result))
  opts := SetQueryOptions(options)

  shards, err := (*kuzzle.Kuzzle)(k.instance).RefreshIndex(C.GoString(index), opts)
  if err != nil {
    Set_shards_result_error(result, err)
    return result
  }

  result.result = goToCShards(shards)

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

//export kuzzle_wrapper_get_server_info
func kuzzle_wrapper_get_server_info(k *C.kuzzle, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))
  opts := SetQueryOptions(options)

  res, err := (*kuzzle.Kuzzle)(k.instance).GetServerInfo(opts)

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

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(k *C.kuzzle, options *C.query_options) *C.json_result {
  result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))
  result.result = (*C._json_object)(C.calloc(1, C.sizeof__json_object))

  var opts types.QueryOptions
  if options != nil {
    opts = SetQueryOptions(options)
  }

  res, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)

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
  completedRequests, _ := json.Marshal(res.CompletedRequests)
  connections, _ := json.Marshal(res.Connections)
  failedRequests, _ := json.Marshal(res.FailedRequests)

  cOnGoing := C.CString(string(ongoing))
  cCompleteRequest := C.CString(string(completedRequests))
  cConnections := C.CString(string(connections))
  cFailedRequests := C.CString(string(failedRequests))

  result.ongoing_requests = C.json_tokener_parse(cOnGoing)
  result.completed_requests = C.json_tokener_parse(cCompleteRequest)
  result.connections = C.json_tokener_parse(cConnections)
  result.failed_requests = C.json_tokener_parse(cFailedRequests)
  result.timestamp = C.ulonglong(res.Timestamp)

  C.free(unsafe.Pointer(cOnGoing))
  C.free(unsafe.Pointer(cCompleteRequest))
  C.free(unsafe.Pointer(cConnections))
  C.free(unsafe.Pointer(cFailedRequests))

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
