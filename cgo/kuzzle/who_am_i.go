package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <stdlib.h>
	#include "kuzzle.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
)

//export kuzzle_wrapper_who_am_i
func kuzzle_wrapper_who_am_i(k *C.Kuzzle) *C.user {
	user := (*C.user)(C.calloc(1, C.sizeof_user))

	res, err := (*kuzzle.Kuzzle)(k.instance).WhoAmI()
	if err != nil {
		Set_user_error(user, err)
		return user
	}

	user.meta = goToCKuzzleMeta(res.Meta)

	buffer := C.CString(string(res.Source))
	user.source = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	user.strategies_length = C.ulong(len(res.Strategies))
	user.strategies = (**C.char)(C.calloc(C.size_t(user.strategies_length), C.sizeof_char_ptr))
	cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(user.strategies))[:len(res.Strategies):len(res.Strategies)]

	for i, substring := range res.Strategies {
		cArray[i] = C.CString(substring)
	}

	user.id = C.CString(res.Id)

	return user
}
