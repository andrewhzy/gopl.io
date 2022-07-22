// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
	// "math"
)

func main() {
	test4()
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
	var r = [...]rune{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, '世'}
	var d interface{} = a
	v, _ := d.(int)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(r))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof(v))
	fmt.Println("------------------")
	fmt.Println(unsafe.Alignof(a))
	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Alignof(c))
	fmt.Println(unsafe.Alignof(r))
	fmt.Println(unsafe.Alignof(d))
	arr := [...]interface{}{a, b, c, r, d, v}
	for _, item := range arr {
		fmt.Printf("size: %d, alien: %d, type: %s\n", unsafe.Sizeof(item), unsafe.Alignof(item), item)
	}
}

func test2() {
	a := new(int)
	fmt.Println(&*a)
	fmt.Println(a)
	*a = 3
	fmt.Println(&*a)
	fmt.Println(a)

	b := 1
	//a = &b
	fmt.Println(&b)
	fmt.Println(b)
	*a = 3
	fmt.Println(&b)
	fmt.Println(unsafe.Alignof(unsafe.Pointer(&b)))
	fmt.Println(uintptr(unsafe.Pointer(&b)))
}

func test3() {
	var r = [...]rune{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, '世'}
	fmt.Println(unsafe.Sizeof(r))
	fmt.Println(len(r))

	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i, r := range s {
		fmt.Printf("%3d    %3q    %3d\n", i, r, r)
	}
}

func test4() {
	// "program" in Japanese katakana
	s := "编程语言"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	fmt.Printf("%s\n", s)
	fmt.Printf("%x\n", []byte(s))
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
}
