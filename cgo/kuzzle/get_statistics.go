package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"unsafe"
	"time"
	"strconv"
)

//export kuzzle_wrapper_get_statistics
func kuzzle_wrapper_get_statistics(result *C.statistics, timestamp C.time_t, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	t, _ := strconv.ParseInt(C.GoString(C.ctime(&timestamp)), 10, 64)
	tm := time.Unix(t, 0)

	res, err := KuzzleInstance.GetStatistics(&tm, opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	ongoing, _ := json.Marshal(res.OngoingRequests)
	completed_requests, _ := json.Marshal(res.CompletedRequests)
	connections, _ := json.Marshal(res.Connections)
	failed_requests, _ := json.Marshal(res.FailedRequests)

	result.ongoing_requests = C.json_tokener_parse(C.CString(string(ongoing)))
	result.completed_requests = C.json_tokener_parse(C.CString(string(completed_requests)))
	result.completed_requests = C.json_tokener_parse(C.CString(string(connections)))
	result.completed_requests = C.json_tokener_parse(C.CString(string(failed_requests)))
	result.timestamp = C.double(res.Timestamp)
}
