package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"regexp"
)

//export kuzzle_wrapper_collection_mcreate_document
func kuzzle_wrapper_collection_mcreate_document(c *C.collection, result *C.kuzzle_search_response, documents **C.document, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).MCreateDocument(goDocuments(documents), opts)
	if err != nil {
		match, _ := regexp.MatchString("Invalid value for the 'ifExist' option: '.*'", err.Error())
		if match {
			return C.int(C.EINVAL)
		}
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return 0
	}

	result.instance = unsafe.Pointer(&res)

	return 0
}


//export kuzzle_wrapper_collection_mcreate_or_replace_document
func kuzzle_wrapper_collection_mcreate_or_replace_document(c *C.collection, result *C.kuzzle_search_response, documents **C.document, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).MCreateOrReplaceDocument(goDocuments(documents), opts)
	if err != nil {
		match, _ := regexp.MatchString("Invalid value for the 'ifExist' option: '.*'", err.Error())
		if match {
			return C.int(C.EINVAL)
		}
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return 0
	}

	result.instance = unsafe.Pointer(&res)

	return 0
}