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
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
	"time"
	"unsafe"
)

//export kuzzle_wrapper_get_statistics
func kuzzle_wrapper_get_statistics(k *C.kuzzle, timestamp C.time_t, options *C.query_options) *C.statistics {
	result := (*C.statistics)(C.calloc(1, C.sizeof_statistics))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

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
