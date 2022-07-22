// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 357.

// Package unsafeptr demonstrates basic use of unsafe.Pointer.
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//!+main
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"

	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb = (*int16)(unsafe.Pointer(tmp))
	*pb = 42
	fmt.Println(x.b) // "42"
	//!-main

	fmt.Printf("%#016x\n", math.Float64bits(1.0))
	fmt.Printf("%#016x\n", uintptr(unsafe.Pointer(&x)))
}

/*
//!+wrong
	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
//!-wrong
*/

//func test(){
//	package math
//	func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }
//	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"
//}

//0x3ff0000000000000
