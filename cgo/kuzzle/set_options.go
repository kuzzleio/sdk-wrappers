package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
 */
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func SetOptions(options *C.query_options) types.QueryOptions {
	opts := types.NewQueryOptions()

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

	out, _ := json.Marshal(opts.GetVolatile())
	vols := make(map[string]interface{})
	json.Unmarshal(out, &vols)
	opts.SetVolatile(vols)

	return opts
}
