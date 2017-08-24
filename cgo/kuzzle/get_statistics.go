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
		opts = SetOptions(options)
	}

	t, _ := strconv.ParseInt(C.GoString(C.ctime(&timestamp)), 10, 64)
	tm := time.Unix(t, 0)

	res, err := KuzzleInstance.GetStatistics(&tm, opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	ongoing, _ := json.Marshal(res.OngoingRequests)
	result.ongoing_requests = C.json_tokener_parse(C.CString(string(ongoing)))
}