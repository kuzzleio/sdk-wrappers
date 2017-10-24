package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

func SetQueryOptions(options *C.query_options) (opts types.QueryOptions) {
	opts = types.NewQueryOptions()

	opts.SetQueuable(bool(options.queuable))
	opts.SetFrom(int(options.from))
	opts.SetSize(int(options.size))

	opts.SetScroll(C.GoString(options.scroll))
	opts.SetScrollId(C.GoString(options.scroll_id))
	opts.SetRefresh(C.GoString(options.refresh))
	opts.SetIfExist(C.GoString(options.if_exist))
	opts.SetRetryOnConflict(int(options.retry_on_conflict))
	opts.SetVolatile(JsonCConvert(options.volatiles).(map[string]interface{}))

	return
}

func SetOptions(options *C.options) (opts types.Options) {
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
	opts.SetHeaders(JsonCConvert(options.headers).(map[string]interface{}))

	return
}

func SetRoomOptions(options *C.room_options) (opts types.RoomOptions) {
	opts = types.NewRoomOptions()

	opts.SetScope(C.GoString(options.scope))
	opts.SetState(C.GoString(options.state))
	opts.SetUser(C.GoString(options.user))

	opts.SetSubscribeToSelf(options.subscribe_to_self == 1)

	if options.volatiles != nil {
		opts.SetVolatile(JsonCConvert(options.volatiles).(map[string]interface{}))
	}

	return
}
