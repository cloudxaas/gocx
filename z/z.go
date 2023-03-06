// Copyright 2023 CloudXaas.
// 
// Licensed under the MIT License.
// The most frequently used hacks which may not be compatible with golang so we put as "z" for the moment.
package z

import (
	"reflect"
	"unsafe"
)

//Credit : Ian Lance Taylor
//https://groups.google.com/g/golang-nuts/c/Zsfk-VMd_fU/m/O1ru4fO-BgAJ
func S2b(s string) (b []byte) {
    if s == "" {
        return nil // or []byte{}
    }
    return (*[0x7fff0000]byte)(unsafe.Pointer(
        (*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
    )[:len(s):len(s)]
}

func B2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

