package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"
	"time"
	"strconv"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

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
