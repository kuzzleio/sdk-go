package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include "kuzzlesdk.h"
*/
import "C"
import "unsafe"

func JsonCType(jobj *C.json_object) C.json_type {
	// Returning the value directly results in a type mismatch
	switch C.json_object_get_type(jobj) {
	case C.json_type_null:
		return C.json_type_null
	case C.json_type_boolean:
		return C.json_type_boolean
	case C.json_type_double:
		return C.json_type_double
	case C.json_type_int:
		return C.json_type_int
	case C.json_type_object:
		return C.json_type_object
	case C.json_type_array:
		return C.json_type_array
	default:
		return C.json_type_null
	}
}

func JsonCConvert(jobj *C.json_object) interface{} {
	if jobj == nil {
		return nil
	}

	jtype := C.json_object_get_type(jobj)

	switch jtype {
	case C.json_type_null:
		return nil
	case C.json_type_boolean:
		return int(C.json_object_get_boolean(jobj)) == 1
	case C.json_type_double:
		return float64(C.json_object_get_double(jobj))
	case C.json_type_int:
		return int(C.json_object_get_int(jobj))
	case C.json_type_string:
		return C.GoString(C.json_object_get_string(jobj))
	case C.json_type_object:
		table := C.json_object_get_object(jobj)
		content := make(map[string]interface{})

		if table == nil {
			return content
		}

		for field, nextField := table.head, table.head; field != nil; field = nextField {
			nextField = field.next

			key := (*C.char)(field.k)
			value := (*C.json_object)(field.v)

			content[C.GoString(key)] = JsonCConvert(value)
		}

		return content
	case C.json_type_array:
		length := int(C.json_object_array_length(jobj))
		content := make([]interface{}, length)

		for i := 0; i < length; i++ {
			content[i] = JsonCConvert(C.json_object_array_get_idx(jobj, C.size_t(i)))
		}

		return content
	}

	return nil
}

//export kuzzle_wrapper_json_new
func kuzzle_wrapper_json_new(jobj *C.json_object) {
	jobj = C.json_object_new_object()
}

//export kuzzle_wrapper_json_put
func kuzzle_wrapper_json_put(jobj *C.json_object, key *C.char, content unsafe.Pointer, kind C.int) {
	if kind == 0 {
		//string
		C.json_object_object_add(jobj, key, C.json_object_new_string((*C.char)(content)))
	} else if kind == 1 {
		//int
		C.json_object_object_add(jobj, key, C.json_object_new_int64((C.int64_t)(*(*C.int)(content))))
	} else if kind == 2 {
		//double
		C.json_object_object_add(jobj, key, C.json_object_new_double(*(*C.double)(content)))
	} else if kind == 3 {
		//bool
		C.json_object_object_add(jobj, key, C.json_object_new_boolean((C.json_bool)(*(*C.uchar)(content))))
	} else if kind == 4 {
		//json_object
		C.json_object_object_add(jobj, key, (*C.json_object)(content))
	}
}

//export kuzzle_wrapper_json_get_string
func kuzzle_wrapper_json_get_string(jobj *C.json_object, key *C.char) *C.char {
	var value *C.json_object
	C.json_object_object_get_ex(jobj, key, &value)

	return C.json_object_get_string(value)
}

//export kuzzle_wrapper_json_get_int
func kuzzle_wrapper_json_get_int(jobj *C.json_object, key *C.char) C.int {
	var value *C.json_object
	C.json_object_object_get_ex(jobj, key, &value)

	return C.int(C.json_object_get_int64(value))
}

//export kuzzle_wrapper_json_get_double
func kuzzle_wrapper_json_get_double(jobj *C.json_object, key *C.char) C.double {
	var value *C.json_object
	C.json_object_object_get_ex(jobj, key, &value)

	return C.json_object_get_double(value)
}

//export kuzzle_wrapper_json_get_bool
func kuzzle_wrapper_json_get_bool(jobj *C.json_object, key *C.char) C.json_bool {
	var value *C.json_object
	C.json_object_object_get_ex(jobj, key, &value)

	return C.json_object_get_boolean(value)
}

//export kuzzle_wrapper_json_get_json_object
func kuzzle_wrapper_json_get_json_object(jobj *C.json_object, key *C.char) *C.json_object {
	var value *C.json_object

	C.json_object_object_get_ex(jobj, key, &value)
	return value
}
