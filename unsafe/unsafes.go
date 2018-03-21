// Copyright © 2018 tzx All rights reserved.

package unsafe

import (
	"reflect"
	"unsafe"
)

// String change byte slice to string
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes change string to byte slice.
func Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// ByteString no copy to change byte slice to string.
func ByteString(b []byte) (s string) {
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ps.Data = pb.Data
	ps.Len = pb.Len
	return
}

// StringByte no copy to change string to byte slice.
func StringByte(s string) (b []byte) {
	pb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ps := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pb.Data = ps.Data
	pb.Len = ps.Len
	pb.Cap = ps.Len
	return
}

// SliceEqual return true if two slices equal.
func SliceEqual(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// SliceEqualBCE return true if two slices equal. Add BCE check before for range loop.
func SliceEqualBCE(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// SliceReflectEqual return true if two slices equal.
func SliceReflectEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
