package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"


//export kuzzle_wrapper_collection_count
func kuzzle_wrapper_new_document(c *C.collection) *C.document {
	return cToGoCollection(c).Document()
}

func kuzzle_wrapper_document_fetch(d *C.document) {
	// TODO
}