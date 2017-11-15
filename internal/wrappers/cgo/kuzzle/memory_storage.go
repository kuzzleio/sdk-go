package main

/*
  #cgo CFLAGS: -I../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include <string.h>
  #include <json-c/json.h>
  #include "kuzzlesdk.h"
  #include "sdk_wrappers_internal.h"

  static void assign_geopos(double (*ptr)[2], int idx, double lon, double lat) {
    ptr[idx][0] = lon;
    ptr[idx][1] = lat;
  }
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_ms_append
func kuzzle_wrapper_ms_append(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Append(
		C.GoString(key),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_bitcount
func kuzzle_wrapper_ms_bitcount(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitcount(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_bitop
func kuzzle_wrapper_ms_bitop(k *C.kuzzle, key *C.char, operation *C.char, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitop(
		C.GoString(key),
		C.GoString(operation),
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_bitpos
func kuzzle_wrapper_ms_bitpos(k *C.kuzzle, key *C.char, bit C.uchar, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Bitpos(
		C.GoString(key),
		int(bit),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_dbsize
func kuzzle_wrapper_ms_dbsize(k *C.kuzzle, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Dbsize(SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_decr
func kuzzle_wrapper_ms_decr(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Decr(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_decrby
func kuzzle_wrapper_ms_decrby(k *C.kuzzle, key *C.char, value C.int, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Decrby(
		C.GoString(key),
		int(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_del
func kuzzle_wrapper_ms_del(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Del(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_exists
func kuzzle_wrapper_ms_exists(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Exists(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_expire
func kuzzle_wrapper_ms_expire(k *C.kuzzle, key *C.char, seconds C.ulong, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Expire(
		C.GoString(key),
		int(seconds),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_expireat
func kuzzle_wrapper_ms_expireat(k *C.kuzzle, key *C.char, ts C.ulonglong, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Expireat(
		C.GoString(key),
		int(ts),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_flushdb
func kuzzle_wrapper_ms_flushdb(k *C.kuzzle, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Flushdb(SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_geoadd
func kuzzle_wrapper_ms_geoadd(k *C.kuzzle, key *C.char, points **C.json_object, plen C.size_t, options *C.query_options) *C.int_result {
	wrapped := (*[1 << 20]*C.json_object)(unsafe.Pointer(points))[:plen:plen]
	gopoints := make([]*types.GeoPoint, int(plen))

	for i, jobj := range wrapped {
		stringified := C.json_object_to_json_string(jobj)
		gobytes := C.GoBytes(unsafe.Pointer(stringified), C.int(C.strlen(stringified)))
		json.Unmarshal(gobytes, gopoints[i])
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Geoadd(
		C.GoString(key),
		gopoints,
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_geodist
func kuzzle_wrapper_ms_geodist(k *C.kuzzle, key *C.char, member1 *C.char, member2 *C.char, options *C.query_options) *C.double_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Geodist(
		C.GoString(key),
		C.GoString(member1),
		C.GoString(member2),
		SetQueryOptions(options))

	return goToCDoubleResult(res, err)
}

//export kuzzle_wrapper_ms_geohash
func kuzzle_wrapper_ms_geohash(k *C.kuzzle, key *C.char, members **C.char, mlen C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Geohash(
		C.GoString(key),
		cToGoStrings(members, mlen),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_geopos
func kuzzle_wrapper_ms_geopos(k *C.kuzzle, key *C.char, members **C.char, mlen C.size_t, options *C.query_options) *C.geopos_result {
	result := (*C.geopos_result)(C.calloc(1, C.sizeof_geopos_result))

	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Geopos(
		C.GoString(key),
		cToGoStrings(members, mlen),
		SetQueryOptions(options))

	if err != nil {
		kuzzleError := err.(*types.KuzzleError)
		result.status = C.int(kuzzleError.Status)
		result.error = C.CString(kuzzleError.Message)

		if len(kuzzleError.Stack) > 0 {
			result.stack = C.CString(kuzzleError.Stack)
		}
		return result
	}

	result.result_length = C.size_t(len(res))
	result.result = (*[2]C.double)(C.calloc(result.result_length, C.sizeof_geopos_arr))

	for i, pos := range res {
		C.assign_geopos(result.result, C.int(i), C.double(pos.Lon), C.double(pos.Lat))
	}

	return result
}

//export kuzzle_wrapper_ms_georadius
func kuzzle_wrapper_ms_georadius(k *C.kuzzle, key *C.char, lon C.double, lat C.double, dist C.double, unit *C.char, options *C.query_options) *C.json_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Georadius(
		C.GoString(key),
		float64(lon),
		float64(lat),
		float64(dist),
		C.GoString(unit),
		SetQueryOptions(options))

	var ires []interface{}
	if err == nil {
		ires = make([]interface{}, len(res))
		for i, d := range res {
			ires[i] = d
		}
	}

	return goToCJsonArrayResult(ires, err)
}

//export kuzzle_wrapper_ms_georadiusbymember
func kuzzle_wrapper_ms_georadiusbymember(k *C.kuzzle, key *C.char, member *C.char, dist C.double, unit *C.char, options *C.query_options) *C.json_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Georadiusbymember(
		C.GoString(key),
		C.GoString(member),
		float64(dist),
		C.GoString(unit),
		SetQueryOptions(options))

	var ires []interface{}
	if err == nil {
		ires = make([]interface{}, len(res))
		for i, d := range res {
			ires[i] = d
		}
	}

	return goToCJsonArrayResult(ires, err)
}

//export kuzzle_wrapper_ms_get
func kuzzle_wrapper_ms_get(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Get(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_getbit
func kuzzle_wrapper_ms_getbit(k *C.kuzzle, key *C.char, offset C.int, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Getbit(
		C.GoString(key),
		int(offset),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_getrange
func kuzzle_wrapper_ms_getrange(k *C.kuzzle, key *C.char, start C.int, end C.int, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Getrange(
		C.GoString(key),
		int(start),
		int(end),
		SetQueryOptions(options))

	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_ms_getset
func kuzzle_wrapper_ms_getset(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Getset(
		C.GoString(key),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_hdel
func kuzzle_wrapper_ms_hdel(k *C.kuzzle, key *C.char, fields **C.char, flen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hdel(
		C.GoString(key),
		cToGoStrings(fields, flen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_hexists
func kuzzle_wrapper_ms_hexists(k *C.kuzzle, key *C.char, field *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hexists(
		C.GoString(key),
		C.GoString(field),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_hget
func kuzzle_wrapper_ms_hget(k *C.kuzzle, key *C.char, field *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hget(
		C.GoString(key),
		C.GoString(field),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_hgetall
func kuzzle_wrapper_ms_hgetall(k *C.kuzzle, key *C.char, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hgetall(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_ms_hincrby
func kuzzle_wrapper_ms_hincrby(k *C.kuzzle, key *C.char, field *C.char, value C.long, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hincrby(
		C.GoString(key),
		C.GoString(field),
		int(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_hincrbyfloat
func kuzzle_wrapper_ms_hincrbyfloat(k *C.kuzzle, key *C.char, field *C.char, value C.double, options *C.query_options) *C.double_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hincrbyfloat(
		C.GoString(key),
		C.GoString(field),
		float64(value),
		SetQueryOptions(options))

	return goToCDoubleResult(res, err)
}

//export kuzzle_wrapper_ms_hkeys
func kuzzle_wrapper_ms_hkeys(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hkeys(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_hlen
func kuzzle_wrapper_ms_hlen(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hlen(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_hmget
func kuzzle_wrapper_ms_hmget(k *C.kuzzle, key *C.char, fields **C.char, flen C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hmget(
		C.GoString(key),
		cToGoStrings(fields, flen),
		SetQueryOptions(options))

	// Ms.Hmget returns a []*string value and goToCStringArrayResult
	// only accept a []string one
	var converted []string

	if err == nil {
		converted = make([]string, len(res), len(res))

		for i, val := range res {
			converted[i] = *val
		}
	}

	return goToCStringArrayResult(converted, err)
}

//export kuzzle_wrapper_ms_hmset
func kuzzle_wrapper_ms_hmset(k *C.kuzzle, key *C.char, entries **C.json_object, elen C.size_t, options *C.query_options) *C.void_result {
	wrapped := (*[1 << 20]*C.json_object)(unsafe.Pointer(entries))[:elen:elen]
	goentries := make([]*types.MsHashField, int(elen))

	for i, jobj := range wrapped {
		stringified := C.json_object_to_json_string(jobj)
		gobytes := C.GoBytes(unsafe.Pointer(stringified), C.int(C.strlen(stringified)))
		json.Unmarshal(gobytes, goentries[i])
	}

	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hmset(
		C.GoString(key),
		goentries,
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_hscan
func kuzzle_wrapper_ms_hscan(k *C.kuzzle, key *C.char, cursor C.int, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hscan(
		C.GoString(key),
		int(cursor),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_ms_hset
func kuzzle_wrapper_ms_hset(k *C.kuzzle, key *C.char, field *C.char, value *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hset(
		C.GoString(key),
		C.GoString(field),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_hsetnx
func kuzzle_wrapper_ms_hsetnx(k *C.kuzzle, key *C.char, field *C.char, value *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hsetnx(
		C.GoString(key),
		C.GoString(field),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_hstrlen
func kuzzle_wrapper_ms_hstrlen(k *C.kuzzle, key *C.char, field *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hstrlen(
		C.GoString(key),
		C.GoString(field),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_hvals
func kuzzle_wrapper_ms_hvals(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Hvals(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_incr
func kuzzle_wrapper_ms_incr(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Incr(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_incrby
func kuzzle_wrapper_ms_incrby(k *C.kuzzle, key *C.char, value C.long, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Incrby(
		C.GoString(key),
		int(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_incrbyfloat
func kuzzle_wrapper_ms_incrbyfloat(k *C.kuzzle, key *C.char, value C.double, options *C.query_options) *C.double_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Incrbyfloat(
		C.GoString(key),
		float64(value),
		SetQueryOptions(options))

	return goToCDoubleResult(res, err)
}

//export kuzzle_wrapper_ms_keys
func kuzzle_wrapper_ms_keys(k *C.kuzzle, pattern *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Keys(
		C.GoString(pattern),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_lindex
func kuzzle_wrapper_ms_lindex(k *C.kuzzle, key *C.char, index C.long, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lindex(
		C.GoString(key),
		int(index),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_linsert
func kuzzle_wrapper_ms_linsert(k *C.kuzzle, key *C.char, position *C.char, pivot *C.char, value *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Linsert(
		C.GoString(key),
		C.GoString(position),
		C.GoString(pivot),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_llen
func kuzzle_wrapper_ms_llen(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Llen(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_lpop
func kuzzle_wrapper_ms_lpop(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lpop(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_lpush
func kuzzle_wrapper_ms_lpush(k *C.kuzzle, key *C.char, values **C.char, vlen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lpush(
		C.GoString(key),
		cToGoStrings(values, vlen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_lpushx
func kuzzle_wrapper_ms_lpushx(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lpushx(
		C.GoString(key),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_lrange
func kuzzle_wrapper_ms_lrange(k *C.kuzzle, key *C.char, start C.long, stop C.long, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lrange(
		C.GoString(key),
		int(start),
		int(stop),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_lrem
func kuzzle_wrapper_ms_lrem(k *C.kuzzle, key *C.char, count C.long, value *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lrem(
		C.GoString(key),
		int(count),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_lset
func kuzzle_wrapper_ms_lset(k *C.kuzzle, key *C.char, index C.long, value *C.char, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Lset(
		C.GoString(key),
		int(index),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_ltrim
func kuzzle_wrapper_ms_ltrim(k *C.kuzzle, key *C.char, start C.long, stop C.long, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Ltrim(
		C.GoString(key),
		int(start),
		int(stop),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_mget
func kuzzle_wrapper_ms_mget(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Mget(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	var converted []string

	if err == nil {
		converted = make([]string, len(res), len(res))

		for i, val := range res {
			converted[i] = *val
		}
	}

	return goToCStringArrayResult(converted, err)
}

//export kuzzle_wrapper_ms_mset
func kuzzle_wrapper_ms_mset(k *C.kuzzle, entries **C.json_object, elen C.size_t, options *C.query_options) *C.void_result {
	wrapped := (*[1 << 20]*C.json_object)(unsafe.Pointer(entries))[:elen:elen]
	goentries := make([]*types.MSKeyValue, int(elen))

	for i, jobj := range wrapped {
		stringified := C.json_object_to_json_string(jobj)
		gobytes := C.GoBytes(unsafe.Pointer(stringified), C.int(C.strlen(stringified)))
		json.Unmarshal(gobytes, goentries[i])
	}

	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Mset(
		goentries,
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_msetnx
func kuzzle_wrapper_ms_msetnx(k *C.kuzzle, entries **C.json_object, elen C.size_t, options *C.query_options) *C.bool_result {
	wrapped := (*[1 << 20]*C.json_object)(unsafe.Pointer(entries))[:elen:elen]
	goentries := make([]*types.MSKeyValue, int(elen))

	for i, jobj := range wrapped {
		stringified := C.json_object_to_json_string(jobj)
		gobytes := C.GoBytes(unsafe.Pointer(stringified), C.int(C.strlen(stringified)))
		json.Unmarshal(gobytes, goentries[i])
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Msetnx(
		goentries,
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_object
func kuzzle_wrapper_ms_object(k *C.kuzzle, key *C.char, subcommand *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Object(
		C.GoString(key),
		C.GoString(subcommand),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_persist
func kuzzle_wrapper_ms_persist(k *C.kuzzle, key *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Persist(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_pexpire
func kuzzle_wrapper_ms_pexpire(k *C.kuzzle, key *C.char, ttl C.ulong, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Pexpire(
		C.GoString(key),
		int(ttl),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_pexpireat
func kuzzle_wrapper_ms_pexpireat(k *C.kuzzle, key *C.char, ts C.ulonglong, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Pexpireat(
		C.GoString(key),
		int(ts),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_pfadd
func kuzzle_wrapper_ms_pfadd(k *C.kuzzle, key *C.char, elements **C.char, elen C.size_t, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Pfadd(
		C.GoString(key),
		cToGoStrings(elements, elen),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_pfcount
func kuzzle_wrapper_ms_pfcount(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Pfcount(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_pfmerge
func kuzzle_wrapper_ms_pfmerge(k *C.kuzzle, key *C.char, sources **C.char, slen C.size_t, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Pfmerge(
		C.GoString(key),
		cToGoStrings(sources, slen),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_ping
func kuzzle_wrapper_ms_ping(k *C.kuzzle, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Ping(
		SetQueryOptions(options))

	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_ms_psetex
func kuzzle_wrapper_ms_psetex(k *C.kuzzle, key *C.char, value *C.char, ttl C.ulong, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Psetex(
		C.GoString(key),
		C.GoString(value),
		int(ttl),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_pttl
func kuzzle_wrapper_ms_pttl(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Pttl(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_randomkey
func kuzzle_wrapper_ms_randomkey(k *C.kuzzle, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Randomkey(
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_rename
func kuzzle_wrapper_ms_rename(k *C.kuzzle, key *C.char, newkey *C.char, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Rename(
		C.GoString(key),
		C.GoString(newkey),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_renamenx
func kuzzle_wrapper_ms_renamenx(k *C.kuzzle, key *C.char, newkey *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Renamenx(
		C.GoString(key),
		C.GoString(newkey),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_rpop
func kuzzle_wrapper_ms_rpop(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Rpop(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_rpoplpush
func kuzzle_wrapper_ms_rpoplpush(k *C.kuzzle, key *C.char, dest *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Rpoplpush(
		C.GoString(key),
		C.GoString(dest),
		SetQueryOptions(options))

	return goToCStringResult(res, err)
}

//export kuzzle_wrapper_ms_rpush
func kuzzle_wrapper_ms_rpush(k *C.kuzzle, key *C.char, values **C.char, vlen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Rpush(
		C.GoString(key),
		cToGoStrings(values, vlen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_rpushx
func kuzzle_wrapper_ms_rpushx(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Rpushx(
		C.GoString(key),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_sadd
func kuzzle_wrapper_ms_sadd(k *C.kuzzle, key *C.char, members **C.char, mlen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sadd(
		C.GoString(key),
		cToGoStrings(members, mlen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_scan
func kuzzle_wrapper_ms_scan(k *C.kuzzle, cursor C.int, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Scan(
		int(cursor),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_ms_scard
func kuzzle_wrapper_ms_scard(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Scard(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_sdiff
func kuzzle_wrapper_ms_sdiff(k *C.kuzzle, key *C.char, keys **C.char, klen C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sdiff(
		C.GoString(key),
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_sdiffstore
func kuzzle_wrapper_ms_sdiffstore(k *C.kuzzle, key *C.char, keys **C.char, klen C.size_t, dest *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sdiffstore(
		C.GoString(key),
		cToGoStrings(keys, klen),
		C.GoString(dest),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_set
func kuzzle_wrapper_ms_set(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Set(
		C.GoString(key),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_setex
func kuzzle_wrapper_ms_setex(k *C.kuzzle, key *C.char, value *C.char, ttl C.ulong, options *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Setex(
		C.GoString(key),
		C.GoString(value),
		int(ttl),
		SetQueryOptions(options))

	return goToCVoidResult(err)
}

//export kuzzle_wrapper_ms_setnx
func kuzzle_wrapper_ms_setnx(k *C.kuzzle, key *C.char, value *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Setnx(
		C.GoString(key),
		C.GoString(value),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_sinter
func kuzzle_wrapper_ms_sinter(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sinter(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_sinterstore
func kuzzle_wrapper_ms_sinterstore(k *C.kuzzle, dest *C.char, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sinterstore(
		C.GoString(dest),
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_sismember
func kuzzle_wrapper_ms_sismember(k *C.kuzzle, key *C.char, member *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sismember(
		C.GoString(key),
		C.GoString(member),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_smembers
func kuzzle_wrapper_ms_smembers(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Smembers(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_smove
func kuzzle_wrapper_ms_smove(k *C.kuzzle, key *C.char, dest *C.char, member *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Smove(
		C.GoString(key),
		C.GoString(dest),
		C.GoString(member),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_ms_sort
func kuzzle_wrapper_ms_sort(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sort(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_spop
func kuzzle_wrapper_ms_spop(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Spop(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_srandmember
func kuzzle_wrapper_ms_srandmember(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Srandmember(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_srem
func kuzzle_wrapper_ms_srem(k *C.kuzzle, key *C.char, members **C.char, mlen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Srem(
		C.GoString(key),
		cToGoStrings(members, mlen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_sscan
func kuzzle_wrapper_ms_sscan(k *C.kuzzle, key *C.char, cursor C.int, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sscan(
		C.GoString(key),
		int(cursor),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_ms_strlen
func kuzzle_wrapper_ms_strlen(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Strlen(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_sunion
func kuzzle_wrapper_ms_sunion(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sunion(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_sunionstore
func kuzzle_wrapper_ms_sunionstore(k *C.kuzzle, dest *C.char, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Sunionstore(
		C.GoString(dest),
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_time
func kuzzle_wrapper_ms_time(k *C.kuzzle, options *C.query_options) *C.int_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Time(
		SetQueryOptions(options))

	return goToCIntArrayResult(res, err)
}

//export kuzzle_wrapper_ms_touch
func kuzzle_wrapper_ms_touch(k *C.kuzzle, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Touch(
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_ttl
func kuzzle_wrapper_ms_ttl(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Ttl(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_type
func kuzzle_wrapper_ms_type(k *C.kuzzle, key *C.char, options *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Type(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_ms_zadd
func kuzzle_wrapper_ms_zadd(k *C.kuzzle, key *C.char, elements **C.json_object, elen C.size_t, options *C.query_options) *C.int_result {
	wrapped := (*[1 << 20]*C.json_object)(unsafe.Pointer(elements))[:elen:elen]
	goelements := make([]*types.MSSortedSet, int(elen))

	for i, jobj := range wrapped {
		stringified := C.json_object_to_json_string(jobj)
		gobytes := C.GoBytes(unsafe.Pointer(stringified), C.int(C.strlen(stringified)))
		json.Unmarshal(gobytes, goelements[i])
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zadd(
		C.GoString(key),
		goelements,
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zcard
func kuzzle_wrapper_ms_zcard(k *C.kuzzle, key *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zcard(
		C.GoString(key),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zcount
func kuzzle_wrapper_ms_zcount(k *C.kuzzle, key *C.char, min C.long, max C.long, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zcount(
		C.GoString(key),
		int(min),
		int(max),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zincrby
func kuzzle_wrapper_ms_zincrby(k *C.kuzzle, key *C.char, member *C.char, incr C.double, options *C.query_options) *C.double_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zincrby(
		C.GoString(key),
		C.GoString(member),
		float64(incr),
		SetQueryOptions(options))

	return goToCDoubleResult(res, err)
}

//export kuzzle_wrapper_ms_zinterstore
func kuzzle_wrapper_ms_zinterstore(k *C.kuzzle, dest *C.char, keys **C.char, klen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zinterstore(
		C.GoString(dest),
		cToGoStrings(keys, klen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zlexcount
func kuzzle_wrapper_ms_zlexcount(k *C.kuzzle, key *C.char, min *C.char, max *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zlexcount(
		C.GoString(key),
		C.GoString(min),
		C.GoString(max),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zrange
func kuzzle_wrapper_ms_zrange(k *C.kuzzle, key *C.char, start C.long, stop C.long, options *C.query_options) *C.json_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrange(
		C.GoString(key),
		int(start),
		int(stop),
		SetQueryOptions(options))

	var converted []interface{}

	if res != nil {
		converted = make([]interface{}, len(res), len(res))

		for i, val := range res {
			converted[i] = *val
		}
	}

	return goToCJsonArrayResult(converted, err)
}

//export kuzzle_wrapper_ms_zrangebylex
func kuzzle_wrapper_ms_zrangebylex(k *C.kuzzle, key *C.char, min *C.char, max *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrangebylex(
		C.GoString(key),
		C.GoString(min),
		C.GoString(max),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_zrangebyscore
func kuzzle_wrapper_ms_zrangebyscore(k *C.kuzzle, key *C.char, min C.double, max C.double, options *C.query_options) *C.json_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrangebyscore(
		C.GoString(key),
		float64(min),
		float64(max),
		SetQueryOptions(options))

	var converted []interface{}

	if res != nil {
		converted = make([]interface{}, len(res), len(res))

		for i, val := range res {
			converted[i] = *val
		}
	}

	return goToCJsonArrayResult(converted, err)
}

//export kuzzle_wrapper_ms_zrank
func kuzzle_wrapper_ms_zrank(k *C.kuzzle, key *C.char, member *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrank(
		C.GoString(key),
		C.GoString(member),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zrem
func kuzzle_wrapper_ms_zrem(k *C.kuzzle, key *C.char, members **C.char, mlen C.size_t, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrem(
		C.GoString(key),
		cToGoStrings(members, mlen),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zremrangebylex
func kuzzle_wrapper_ms_zremrangebylex(k *C.kuzzle, key *C.char, min *C.char, max *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zremrangebylex(
		C.GoString(key),
		C.GoString(min),
		C.GoString(max),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zremrangebyrank
func kuzzle_wrapper_ms_zremrangebyrank(k *C.kuzzle, key *C.char, min C.long, max C.long, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zremrangebyrank(
		C.GoString(key),
		int(min),
		int(max),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zremrangebyscore
func kuzzle_wrapper_ms_zremrangebyscore(k *C.kuzzle, key *C.char, min C.double, max C.double, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zremrangebyscore(
		C.GoString(key),
		float64(min),
		float64(max),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zrevrange
func kuzzle_wrapper_ms_zrevrange(k *C.kuzzle, key *C.char, start C.long, stop C.long, options *C.query_options) *C.json_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrevrange(
		C.GoString(key),
		int(start),
		int(stop),
		SetQueryOptions(options))

	var converted []interface{}

	if res != nil {
		converted = make([]interface{}, len(res), len(res))

		for i, val := range res {
			converted[i] = *val
		}
	}

	return goToCJsonArrayResult(converted, err)
}

//export kuzzle_wrapper_ms_zrevrangebylex
func kuzzle_wrapper_ms_zrevrangebylex(k *C.kuzzle, key *C.char, min *C.char, max *C.char, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrevrangebylex(
		C.GoString(key),
		C.GoString(min),
		C.GoString(max),
		SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_ms_zrevrangebyscore
func kuzzle_wrapper_ms_zrevrangebyscore(k *C.kuzzle, key *C.char, min C.double, max C.double, options *C.query_options) *C.json_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrevrangebyscore(
		C.GoString(key),
		float64(min),
		float64(max),
		SetQueryOptions(options))

	var converted []interface{}

	if res != nil {
		converted = make([]interface{}, len(res), len(res))

		for i, val := range res {
			converted[i] = *val
		}
	}

	return goToCJsonArrayResult(converted, err)
}

//export kuzzle_wrapper_ms_zrevrank
func kuzzle_wrapper_ms_zrevrank(k *C.kuzzle, key *C.char, member *C.char, options *C.query_options) *C.int_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zrevrank(
		C.GoString(key),
		C.GoString(member),
		SetQueryOptions(options))

	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_ms_zscan
func kuzzle_wrapper_ms_zscan(k *C.kuzzle, key *C.char, cursor C.int, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zscan(
		C.GoString(key),
		int(cursor),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_ms_zscore
func kuzzle_wrapper_ms_zscore(k *C.kuzzle, key *C.char, member *C.char, options *C.query_options) *C.double_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).MemoryStorage.Zscore(
		C.GoString(key),
		C.GoString(member),
		SetQueryOptions(options))

	return goToCDoubleResult(res, err)
}
