package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
 */
import "C"

type JsonParser struct {
	content map[string]interface{}
}

func (parser *JsonParser) get_json_value(key string, jobj *C.json_object, content map[string]interface{}) {
	jtype := C.json_object_get_type(jobj)

	switch jtype {
	case C.json_type_boolean:
		if int(C.json_object_get_boolean(jobj)) == 1 {
			content[key] = true
		} else {
			content[key] = false
		}
		break
	case C.json_type_string:
		content[key] = C.GoString(C.json_object_get_string(jobj))
		break
	case C.json_type_double:
		content[key] = float64(C.json_object_get_double(jobj))
		break
	case C.json_type_int:
		content[key] = int(C.json_object_get_int(jobj))
		break
	}
}

func (parser *JsonParser) Parse(jobj *C.json_object) {
	if parser.content == nil {
		parser.content = make(map[string]interface{})
	}
	parser.parse_cjson(jobj, parser.content)
}

func (parser *JsonParser) parse_cjson(jobj *C.json_object, content map[string]interface{}) {
	if jobj == nil {
		return
	}

	for entry, entry_next := C.json_object_get_object(jobj).head, C.json_object_get_object(jobj).head; entry != nil; entry = entry_next {
		entry_next = entry.next

		key := (*C.char)(entry.k)
		value := (*C.json_object)(entry.v)

		jtype := C.json_object_get_type(value)

		switch jtype {
		case C.json_type_boolean, C.json_type_double, C.json_type_int, C.json_type_string:
			parser.get_json_value(C.GoString(key), value, content)
			break
		case C.json_type_object:
			C.json_object_object_get_ex(jobj, key, &jobj)
			content[C.GoString(key)] = make(map[string]interface{})
			parser.parse_cjson(jobj, content[C.GoString(key)].(map[string]interface{}))
			break
		case C.json_type_array:

			break
		}
	}
}

func (parser *JsonParser) parseArray(jobj *C.json_object, key *C.char, content map[string]interface{}) {
	jarray := jobj

	if key != nil {
		C.json_object_object_get_ex(jobj, key, &jarray)
	}

	arraylen := C.json_object_array_length(jarray)
	var jvalue *C.json_object

	for i := 0; i < int(arraylen); i++ {
		jvalue = C.json_object_array_get_idx(jarray, C.int(i))
		jtype := C.json_object_get_type(jvalue)
		if jtype == C.json_type_array {
			parser.parseArray(jvalue, nil, content)
		} else if jtype != C.json_type_object {
			parser.get_json_value(C.GoString(key), jvalue, content)
		} else {
			parser.parse_cjson(jvalue, content)
		}
	}
}

func (parser JsonParser) GetContent() map[string]interface{} {
	return parser.content
}