// Copyright 2023 CloudXaaS.
// 
// Licensed under the MIT License.
// The most frequently used hacks which may not be compatible with future golang. benchmark it, it's faster
package cx

import (
//	"reflect"
	"unsafe"
)

/* old
//Credit : Ian Lance Taylor
//https://groups.google.com/g/golang-nuts/c/Zsfk-VMd_fU/m/O1ru4fO-BgAJ
func S2b(s string) (b []byte) {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh.Data = sh.Data
    bh.Cap = sh.Len
    bh.Len = sh.Len
    return b
}
*/

func S2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

/*
func B2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
*/

func B2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
