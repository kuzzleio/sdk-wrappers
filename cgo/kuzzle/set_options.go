package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

func SetQueryOptions(options *C.query_options) (opts types.QueryOptions) {
	opts = types.NewQueryOptions()

	if options.queuable == 0 {
		opts.SetQueuable(false)
	} else {
		opts.SetQueuable(true)
	}
	opts.SetFrom(int(options.from))
	opts.SetSize(int(options.size))
	opts.SetScroll(C.GoString(&options.scroll[0]))
	opts.SetScrollId(C.GoString(&options.scrollId[0]))
	opts.SetRefresh(C.GoString(&options.refresh[0]))
	opts.SetIfExist(C.GoString(&options.ifExist[0]))
	opts.SetRetryOnConflict(int(options.retryOnConflict))
	opts.SetVolatile(JsonCConvert(options.volatiles).(map[string]interface{}))

	return
}

func SetOptions(options *C.Options) (opts types.Options) {
	opts = types.NewOptions()

	opts.SetQueueTTL(time.Duration(int(options.queue_ttl)))
	opts.SetQueueMaxSize(int(options.queue_max_size))
	opts.SetOfflineMode(int(options.offline_mode))

	var autoQueue bool
	if options.auto_queue == 1 {
		autoQueue = true
	}
	opts.SetAutoQueue(autoQueue)

	var autoReconnect bool
	if options.auto_reconnect == 1 {
		autoReconnect = true
	}
	opts.SetAutoReconnect(autoReconnect)

	var autoReplay bool
	if options.auto_replay == 1 {
		autoReplay = true
	}
	opts.SetAutoReplay(autoReplay)

	var autoResub bool
	if options.auto_resubscribe == 1 {
		autoResub = true
	}
	opts.SetAutoResubscribe(autoResub)

	opts.SetReconnectionDelay(time.Duration(int(options.reconnection_delay)))
	opts.SetReplayInterval(time.Duration(int(options.replay_interval)))
	opts.SetConnect(int(options.connect))
	opts.SetRefresh(C.GoString(&options.refresh[0]))
	opts.SetDefaultIndex(C.GoString(&options.default_index[0]))
	opts.SetHeaders(JsonCConvert(options.headers).(map[string]interface{}))

	return
}

func SetRoomOptions(options *C.room_options) (opts types.RoomOptions) {
	opts = types.NewRoomOptions()

	if options.scope != nil {
		opts.SetScope(C.GoString(&options.scope[0]))
	}
	if options.state != nil {
		opts.SetState(C.GoString(&options.state[0]))
	}
	if options.user != nil {
		opts.SetUser(C.GoString(&options.user[0]))
	}

	opts.SetSubscribeToSelf(options.subscribe_to_self == 1)

	if options.volatiles != nil {
		opts.SetVolatile(JsonCConvert(options.volatiles).(map[string]interface{}))
	}

	return
}
