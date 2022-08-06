package ch4

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func add(x int, y int) int { return x + y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x, _ int) int   { return x }
func zero(int, int) int    { return 0 }
func Test0(t *testing.T) {
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func Test1(t *testing.T) {
	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3))     // "3"
	fmt.Printf("%T\n", f) // "func(int) int"

}

func Test2(t *testing.T) {
	var f func(int) int
	f(3) // panic: call of nil function
}

func Test3(t *testing.T) {
	add1 := func(r rune) rune { return r + 1 }
	fmt.Println(strings.Map(add1, "HAL-9000"))
	// "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))   // "WNT"
	fmt.Println(strings.Map(add1, "Admix")) // "Benjy"
}

func Test4(t *testing.T) {
	nums := make([]int, 8)
	for i := range nums {
		nums[i] = i
	}
	var fns []func()
	for i := 0; i < len(nums); i++ {
		//fmt.Println(nums[i])

		n := nums[i]
		fns = append(fns, func() {
			fmt.Println(n)
		})
	}
	for _, f := range fns {
		f()
	}
}

func Test5(t *testing.T) {

	nums := make(map[int]int, 8)
	for i := 0; i < 8; i++ {
		nums[i] = i
	}
	var fns []func()
	for i := 0; i < 8; i++ {
		fmt.Println(nums[i])

		n := nums[i]
		fns = append(fns, func() {
			fmt.Println(n)
		})
	}
	for _, f := range fns {
		f()
	}
}

func f(...int) {}
func g([]int)  {}
func Test6(t *testing.T) {
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
func Test7(t *testing.T) {
	fmt.Println(double(9))
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
func Test8(t *testing.T) {
	fmt.Println(triple(4))
}

func Test9(t *testing.T) {

}

func Test10(t *testing.T) {

}

func Test11(t *testing.T) {

}

func Test12(t *testing.T) {

}

func Test13(t *testing.T) {

}

func Test14(t *testing.T) {

}

func Test15(t *testing.T) {

}

func Test16(t *testing.T) {

}

func Test17(t *testing.T) {

}

func Test18(t *testing.T) {

}

func Test19(t *testing.T) {

	type s1 struct {
		i1 int
		j  int8
	}

	type s2 struct {
		k  s1
		k1 s1
		l  bool
	}
	var a = s1{1, 1}
	var b = s2{a, a, true}
	var c = [...]int{}
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
}
