package main

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_new_options
func kuzzle_wrapper_new_options() *C.Options {
	copts := (*C.Options)(C.calloc(1, C.sizeof_Options))
	opts := types.NewOptions()

	copts.queue_ttl = C.uint(opts.GetQueueTTL())
	copts.queue_max_size = C.ulong(opts.GetQueueMaxSize())
	copts.offline_mode = C.uchar(opts.GetOfflineMode())

	if opts.GetAutoQueue() {
		copts.auto_queue = 1
	}

	if opts.GetAutoReconnect() {
		copts.auto_reconnect = 1
	}

	if opts.GetAutoReplay() {
		copts.auto_replay = 1
	}

	if opts.GetAutoResubscribe() {
		copts.auto_resubscribe = 1
	}

	copts.reconnection_delay = C.ulong(opts.GetReconnectionDelay())
	copts.replay_interval = C.ulong(opts.GetReplayInterval())

	if opts.GetConnect() == 1 {
		copts.connect = C.MANUAL
	} else {
		copts.connect = C.AUTO
	}

	refresh := opts.GetRefresh()
	if len(refresh) > 0 {
		copts.refresh = C.CString(refresh)
	}

	defaultIndex := opts.GetDefaultIndex()
	if len(defaultIndex) > 0 {
		copts.default_index = C.CString(defaultIndex)
	}	

	r, _ := json.Marshal(opts.GetHeaders())
	buffer := C.CString(string(r))
	copts.headers = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return copts
}
