package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"fmt"
	"encoding/json"
)

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(result *C.ack_response, index *C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
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
	}

	var jobj *C.json_object = C.json_tokener_parse(C.CString("{\"foo\":\"bar\", \"az\":{\"second\": \"yihi\"}}"))
	jp := JsonParser{}
	jp.Parse(jobj)

	type ss struct{second string `json:"second"`}
	s := struct{foo string `json:"foo"`; az ss `json:"az"`}{"bar", ss{"yihi"}}

	fmt.Printf("%v\n", jp.GetContent())
	j, _ := json.Marshal(jp.GetContent())
	fmt.Printf("%s", j)

	json.Unmarshal(j, &s)
	fmt.Printf("%v\n", s.az.second)

	//C.json_parse(jobj)

	res, err := KuzzleInstance.CreateIndex(C.GoString(index), opts)
	if err != nil && err.Error() == "Collection.createIndex: index required" {
		return C.int(C.EINVAL)
	}

	var ack, shardsAck C.uint

	if res.Acknowledged {
		ack = 1
	} else {
		ack = 0
	}

	if res.ShardsAcknowledged {
		shardsAck = 1
	} else {
		shardsAck = 0
	}

	*result = C.ack_response{
		acknowledged:       ack,
		shardsAcknowledged: shardsAck,
	}

	return 0
}
