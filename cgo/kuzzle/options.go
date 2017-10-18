package main

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_new_options
func kuzzle_wrapper_new_options(o *C.Options) {
	opts := types.NewOptions()
	copts := C.Options{}

	copts.queue_ttl = C.double(opts.GetQueueTTL())
	copts.queue_max_size = C.int(opts.GetQueueMaxSize())
	copts.offline_mode = C.int(opts.GetOfflineMode())

	var auto_queue uint
	if opts.GetAutoQueue() {
		auto_queue = 1
	}
	copts.auto_queue = C.uint(auto_queue)

	var auto_reconnect uint
	if opts.GetAutoReconnect() {
		auto_reconnect = 1
	}
	copts.auto_reconnect = C.uint(auto_reconnect)

	var auto_replay uint
	if opts.GetAutoReplay() {
		auto_replay = 1
	}
	copts.auto_replay = C.uint(auto_replay)

	var auto_resubscribe uint
	if opts.GetAutoResubscribe() {
		auto_resubscribe = 1
	}
	copts.auto_resubscribe = C.uint(auto_resubscribe)

	copts.reconnection_delay = C.double(opts.GetReconnectionDelay())
	copts.replay_interval = C.double(opts.GetReplayInterval())
	var mode uint32
	if opts.GetConnect() == 1 {
		mode = uint32(C.MANUAL)
	} else {
		mode = uint32(C.AUTO)
	}
	copts.connect = mode

	// TODO Must be freed in C
	copts.refresh = *(*[64]C.char)(unsafe.Pointer(C.CString(opts.GetRefresh())))
	// TODO Must be freed in C
	copts.default_index = *(*[128]C.char)(unsafe.Pointer(C.CString(opts.GetDefaultIndex())))

	r, _ := json.Marshal(opts.GetHeaders())
	headersString := C.CString(string(r))
	defer C.free(unsafe.Pointer(headersString))
	copts.headers = C.json_tokener_parse(headersString)

	*o = copts
}
