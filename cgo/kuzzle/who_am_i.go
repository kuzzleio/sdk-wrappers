package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"unsafe"
)

//export kuzzle_wrapper_who_am_i
func kuzzle_wrapper_who_am_i(user *C.user) {
	res, err := KuzzleInstance.WhoAmI()
	if err != nil {
		user.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
	}

	var meta C.kuzzle_meta
	var active C.uint
	meta.author = *(*[512]C.char)(unsafe.Pointer(C.CString(res.Meta.Author)))
	meta.created_at = C.int(res.Meta.CreatedAt)
	meta.updated_at = C.int(res.Meta.UpdatedAt)
	meta.updater = *(*[512]C.char)(unsafe.Pointer(C.CString(res.Meta.Updater)))

	if res.Meta.Active {
		active = 1
	} else {
		active = 0
	}
	meta.active = active
	meta.deleted_at = C.int(res.Meta.DeletedAt)

	var source *C.json_object
	source = C.json_tokener_parse(C.CString(string(res.Source)))

	cArray := C.malloc(C.size_t(len(res.Strategies)) * C.size_t(unsafe.Sizeof(uintptr(0))))
	a := (*[1<<30 - 1]*C.char)(cArray)
	idx := 0
	for _, substring := range res.Strategies {
		a[idx] = C.CString(substring)
		idx += 1
	}
	a[idx] = nil

	user.id = *(*[512]C.char)(unsafe.Pointer(C.CString(res.Id)))
	user.source = source
	user.meta = &meta
	user.strategies = (**C.char)(cArray)
}
