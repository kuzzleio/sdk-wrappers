package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>

	static int sizeArray(char** arr) {
		int i = 0;

		if (!arr || !arr[0])
			return 0;
		while (arr[i])
			i++;

		return i;
	}
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_query
func kuzzle_wrapper_query(k *C.Kuzzle, result *C.kuzzle_response, request *C.kuzzle_request, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	req := types.KuzzleRequest{
		RequestId:  C.GoString(&request.request_id[0]),
		Controller: C.GoString(&request.controller[0]),
		Action:     C.GoString(&request.action[0]),
		Index:      C.GoString(&request.index[0]),
		Collection: C.GoString(&request.collection[0]),
		Id:         C.GoString(&request.id[0]),
		From:       int(request.from),
		Size:       int(request.size),
		Scroll:     C.GoString(&request.scroll[0]),
		ScrollId:   C.GoString(&request.scroll_id[0]),
		Strategy:   C.GoString(&request.strategy[0]),
		ExpiresIn:  int(request.expires_in),
		Scope:      C.GoString(&request.scope[0]),
		State:      C.GoString(&request.state[0]),
		User:       C.GoString(&request.user[0]),
		Stop:       int(request.stop),
		End:        int(request.end),
		Bit:        int(request.bit),
		Member:     C.GoString(&request.member[0]),
		Member1:    C.GoString(&request.member1[0]),
		Member2:    C.GoString(&request.member2[0]),
		Lon:        float64(request.lon),
		Lat:        float64(request.lat),
		Distance:   float64(request.distance),
		Unit:       C.GoString(&request.unit[0]),
		Offset:     int(request.offset),
		Field:      C.GoString(&request.field[0]),
		Subcommand: C.GoString(&request.subcommand[0]),
		Pattern:    C.GoString(&request.pattern[0]),
		Idx:        int(request.idx),
		Min:        C.GoString(&request.min[0]),
		Max:        C.GoString(&request.max[0]),
		Limit:      C.GoString(&request.limit[0]),
		Count:      int(request.count),
		Match:      C.GoString(&request.match[0]),
	}

	jp := JsonParser{}

	jp.Parse(request.body)
	req.Body = jp.GetContent()

	jp.Parse(request.volatiles)
	req.Volatile = jp.GetContent()

	start := int(request.start)
	req.Start = start

	cursor := int(request.cursor)
	req.Cursor = cursor

	req.Members = goStrings(request.members)
	req.Keys = goStrings(request.keys)
	req.Fields = goStrings(request.fields)

	resC := make(chan *types.KuzzleResponse)
	(*kuzzle.Kuzzle)(k.instance).Query(&req, opts, resC)

	res := <-resC

	if res.Error.Message != "" {
		result.error = ToCString_2048(res.Error.Message)
		return
	}

	result.request_id = ToCString_36(res.RequestId)
	result.room_id = ToCString_36(res.RoomId)
	result.channel = ToCString_128(res.Channel)
	r, _ := json.Marshal(res)
	result.result = C.json_tokener_parse(C.CString(string(r)))
}

// Helper to convert a C char** to a go array of string
func goStrings(argv **C.char) []string {
	length := C.sizeArray(argv)
	if length == 0 {
		return nil
	}
	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
	gostrings := make([]string, length)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
}
