package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
)

// TODO check if it is still legit - refactor
func cToGoSearchFilters(searchFilters *C.search_filters) *types.SearchFilters {
	return &types.SearchFilters{
		Query: JsonCConvert(searchFilters.query),
		Sort: JsonCConvert(searchFilters.sort).([]interface{}),
		Aggregations: JsonCConvert(searchFilters.aggregations),
		SearchAfter: JsonCConvert(searchFilters.search_after).([]interface{}),
	}
}
