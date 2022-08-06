// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import (
	"fmt"
	"unsafe"
	// "math"
)

func main() {
	m := make(map[string]string)
	m[""] = "hello world"
	fmt.Println(m[""])

	name := "世界"
	fmt.Printf("Hello, %s\n", name)
	a := -1 << 63
	if -a > 0 {
		fmt.Println(a)
		fmt.Println(-a)
	}
	test2()
}

type s1 struct {
	i1 int
	j  int8
}

type s2 struct {
	k  s1
	k1 s1
	l  bool
}

func test() {

	var a = s1{1, 1}
	var b = s2{a, a, true}
	var c = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))

	fmt.Println(unsafe.Alignof(a))
	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Alignof(c))
}

func test2() {
	a := new(int)
	fmt.Println(*a)
	*a = 3
	fmt.Println(*a)
}

//!-
