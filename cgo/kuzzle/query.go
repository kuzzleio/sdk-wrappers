package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_query
func kuzzle_wrapper_query(k *C.kuzzle, request *C.kuzzle_request, options *C.query_options) *C.kuzzle_response {
	result := (*C.kuzzle_response)(C.calloc(1, C.sizeof_kuzzle_response))

	if result == nil {
		return result
	}

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	req := types.KuzzleRequest{
		RequestId:  C.GoString(request.request_id),
		Controller: C.GoString(request.controller),
		Action:     C.GoString(request.action),
		Index:      C.GoString(request.index),
		Collection: C.GoString(request.collection),
		Id:         C.GoString(request.id),
		From:       int(request.from),
		Size:       int(request.size),
		Scroll:     C.GoString(request.scroll),
		ScrollId:   C.GoString(request.scroll_id),
		Strategy:   C.GoString(request.strategy),
		ExpiresIn:  int(request.expires_in),
		Scope:      C.GoString(request.scope),
		State:      C.GoString(request.state),
		User:       C.GoString(request.user),
		Start:      int(request.start),
		Stop:       int(request.stop),
		End:        int(request.end),
		Bit:        int(request.bit),
		Member:     C.GoString(request.member),
		Member1:    C.GoString(request.member1),
		Member2:    C.GoString(request.member2),
		Lon:        float64(request.lon),
		Lat:        float64(request.lat),
		Distance:   float64(request.distance),
		Unit:       C.GoString(request.unit),
		Cursor:     int(request.cursor),
		Offset:     int(request.offset),
		Field:      C.GoString(request.field),
		Subcommand: C.GoString(request.subcommand),
		Pattern:    C.GoString(request.pattern),
		Idx:        int(request.idx),
		Min:        C.GoString(request.min),
		Max:        C.GoString(request.max),
		Limit:      C.GoString(request.limit),
		Count:      int(request.count),
		Match:      C.GoString(request.match),
	}

	jp := JsonParser{}

	if request.body != nil {
		jp.Parse(request.body)
		req.Body = jp.GetContent()
	}

	if request.volatiles != nil {
		jp.Parse(request.volatiles)
		req.Volatile = jp.GetContent()
	}

	req.Members = goStrings(request.members, request.members_length)
	req.Keys = goStrings(request.keys, request.keys_length)
	req.Fields = goStrings(request.fields, request.fields_length)

	resC := make(chan *types.KuzzleResponse)
	(*kuzzle.Kuzzle)(k.instance).Query(&req, opts, resC)

	res := <-resC

	if res.Error != nil {
		Set_kuzzle_response_error(result, res.Error)
		return result
	}

	result.request_id = C.CString(res.RequestId)

	if len(res.RoomId) > 0 {
		result.room_id = C.CString(res.RoomId)
	}

	if len(res.Channel) > 0 {
		result.channel = C.CString(res.Channel)
	}

	r, _ := json.Marshal(res)
	buffer := C.CString(string(r))
	result.result = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return result
}

// convert a C char** to a go array of string
func goStrings(arr **C.char, len C.uint) []string {
	if len == 0 {
		return nil
	}

	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(arr))[:len:len]
	gostrings := make([]string, len)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}

	return gostrings
}
