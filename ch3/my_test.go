package ch3

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
	"unicode/utf8"
)

func Test0(t *testing.T) {
	f := 1e100  // a float64
	i := int(f) // result is implementationdependent
	fmt.Println(i)
	fmt.Println(f / 0)
	fmt.Println(int(f))
	f = float64(1 << 100)
	fmt.Println(int(f))
	//fmt.Println(int(float64(1 << 63))) // error: cannot convert float64(1 << 63) (constant 9223372036854775808 of type float64) to type int
}

func Test1(t *testing.T) {
	for x := 0; x < 13; x++ {
		fmt.Printf("x = %2d eA = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func Test2(t *testing.T) {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

}

func Test3(t *testing.T) {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 0 	+Inf	Inf	NaN	"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}

func Test4(t *testing.T) {
	s := "hello 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

func Test5(t *testing.T) {
	s := "hello 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

func Test6(t *testing.T) {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))             // "123 123"
	fmt.Println(strconv.FormatInt(int64(2), 2)) // "1111011"

	fmt.Println(strconv.Atoi("123"))
	fmt.Println(strconv.ParseInt("123", 10, 64)) // base 10, up to 64 bits
}

func Test7(t *testing.T) {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"
}

func Test8(t *testing.T) {

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

}
