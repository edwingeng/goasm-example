//+build !noasm
//+build !appengine

// You must write this file manually.

package main

import (
	"unsafe"
)

//go:noescape
func _AddTwoInts(a int, b int, c unsafe.Pointer)

func AddTwoInts(a, b int, c *int) {
	_AddTwoInts(a, b, unsafe.Pointer(c))
}

//go:noescape
func _SqrtAvx256(in, out unsafe.Pointer)

func SqrtAvx256(in, out []float32) {
	_SqrtAvx256(unsafe.Pointer(&in[0]), unsafe.Pointer(&out[0]))
}
