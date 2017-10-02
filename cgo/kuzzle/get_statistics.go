package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
	"time"
)

//export kuzzle_wrapper_get_statistics
func kuzzle_wrapper_get_statistics(k *C.Kuzzle, result *C.statistics, timestamp C.time_t, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	t, _ := strconv.ParseInt(C.GoString(C.ctime(&timestamp)), 10, 64)
	tm := time.Unix(t, 0)

	res, err := (*kuzzle.Kuzzle)(k.instance).GetStatistics(&tm, opts)
	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	ongoing, _ := json.Marshal(res.OngoingRequests)
	completedRequests, _ := json.Marshal(res.CompletedRequests)
	connections, _ := json.Marshal(res.Connections)
	failedRequests, _ := json.Marshal(res.FailedRequests)

	result.ongoing_requests = C.json_tokener_parse(C.CString(string(ongoing)))
	result.completed_requests = C.json_tokener_parse(C.CString(string(completedRequests)))
	result.completed_requests = C.json_tokener_parse(C.CString(string(connections)))
	result.completed_requests = C.json_tokener_parse(C.CString(string(failedRequests)))
	result.timestamp = C.double(res.Timestamp)
}
