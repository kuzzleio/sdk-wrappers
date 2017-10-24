package main

import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"encoding/json"
	"time"
)

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"

//export kuzzle_wrapper_new_options
func kuzzle_wrapper_new_options() *C.options {
	copts := (*C.options)(C.calloc(1, C.sizeof_options))
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

	default_index := opts.GetDefaultIndex()
	if len(default_index) > 0 {
		copts.default_index = C.CString(default_index)
	}	

	r, _ := json.Marshal(opts.GetHeaders())
	buffer := C.CString(string(r))
	copts.headers = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return copts
}

func SetQueryOptions(options *C.query_options) (opts types.QueryOptions) {
	if options == nil {
		return
	}

	opts = types.NewQueryOptions()

	opts.SetQueuable(bool(options.queuable))
	opts.SetFrom(int(options.from))
	opts.SetSize(int(options.size))
	opts.SetScroll(C.GoString(options.scroll))
	opts.SetScrollId(C.GoString(options.scroll_id))
	opts.SetRefresh(C.GoString(options.refresh))
	opts.SetIfExist(C.GoString(options.if_exist))
	opts.SetRetryOnConflict(int(options.retry_on_conflict))

	out, _ := json.Marshal(opts.GetVolatile())
	vols := make(map[string]interface{})
	json.Unmarshal(out, &vols)
	opts.SetVolatile(vols)

	return
}

func SetOptions(options *C.options) (opts types.Options) {
	if options == nil {
		return
	}

	opts = types.NewOptions()

	opts.SetQueueTTL(time.Duration(uint16(options.queue_ttl)))
	opts.SetQueueMaxSize(int(options.queue_max_size))
	opts.SetOfflineMode(int(options.offline_mode))

	opts.SetAutoQueue(options.auto_queue != 0)
	opts.SetAutoReconnect(options.auto_reconnect != 0)
	opts.SetAutoReplay(options.auto_replay != 0)
	opts.SetAutoResubscribe(options.auto_resubscribe != 0)
	opts.SetReconnectionDelay(time.Duration(int(options.reconnection_delay)))
	opts.SetReplayInterval(time.Duration(int(options.replay_interval)))
	opts.SetConnect(int(options.connect))
	opts.SetRefresh(C.GoString(options.refresh))
	opts.SetDefaultIndex(C.GoString(options.default_index))

	p := JsonParser{}
	p.Parse(options.headers)

	opts.SetHeaders(p.GetContent())

	return
}
