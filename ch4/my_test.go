package ch4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"gopl.io/ch4/nonempty"
)

func Test0(t *testing.T) {
	q := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", q) // "[3]int"
}

func Test1(t *testing.T) {
	q := [3]int{1, 2, 3}
	q = [3]int{1, 2}
	fmt.Println(q)
}

func Test2(t *testing.T) {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	d := [3]int{1, 2}
	fmt.Println(d)
	//fmt.Println(a == d)// compile error: cannot compare [2]int == [3]int
}

func Test3(t *testing.T) {
	var ints []int
	ints = nil
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	fmt.Println(unsafe.Sizeof(ints))
	ints = []int{}
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	fmt.Println(unsafe.Sizeof(ints))
	ints = make([]int, 3)[1:1]
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	fmt.Println(unsafe.Sizeof(ints))

	var intsp *[]int
	intsp = new([]int)
	fmt.Println(len(*intsp))
	fmt.Println(cap(*intsp))
	fmt.Println(unsafe.Sizeof(intsp))
	fmt.Println(unsafe.Sizeof(*intsp))

}

func Test4(t *testing.T) {
	//var ip []
	//ip = ""
	s := ""
	ip := []byte(s)
	ip = make([]byte, 4, 8)
	fmt.Println(len(ip))
	fmt.Println(cap(ip))
	fmt.Println(unsafe.Sizeof(ip))

	is := [1]int{1}
	bs := is[:]
	fmt.Println(len(bs))
	fmt.Println(cap(bs))
	fmt.Println(unsafe.Sizeof(bs))
	//ip = []int{}
	//fmt.Println(len(ip))
	//fmt.Println(unsafe.Sizeof(ip))
	//ip = make([]int, 3)[1:1]
	//fmt.Println(len(ip))
	//fmt.Println(unsafe.Sizeof(ip))
}

func Test5(t *testing.T) {
	var runes []rune
	for _, r := range "Hello, BF" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'B' 'F']"

	fmt.Println(len(runes))
	fmt.Println(cap(runes))
	fmt.Println(unsafe.Sizeof(runes))
}

func Test6(t *testing.T) {
	println("jfkdls")

}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
func Test7(t *testing.T) {
	s := []string{"fjl", "fjsdkl", "1", "2", "3"}
	Add(s)
	fmt.Println(Count(s))
}

// Exercise 4.8
func Test8(t *testing.T) {
	counts := make(map[rune]int)           // counts of Unicode characters
	countsCategory := make(map[string]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int        // count of lengths of UTF-8 encodings
	invalid := 0                           // count of invalid UTF-8 characters

	file, _ := os.Open("my_test.go")
	in := bufio.NewReader(file)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			countsCategory["letter"]++
		}
		if unicode.IsNumber(r) {
			countsCategory["number"]++
		}
		if unicode.IsControl(r) {
			countsCategory["control_code"]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	for c, n := range countsCategory {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}

// Exercise 4.9
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

func TestNonempty0(t *testing.T) {
	fmt.Printf("%q\n", nonempty.Nonempty0(strs))
	fmt.Printf("%q\n", strs)
}

func TestNonempty1(t *testing.T) {
	fmt.Printf("%q\n", nonempty.Nonempty1(strs))
	fmt.Printf("%q\n", strs)
}

func TestNonempty2(t *testing.T) {
	fmt.Printf("%q\n", nonempty.Nonempty2(strs))
	fmt.Printf("%q\n", strs)
}

func TestNonempty3(t *testing.T) {
	fmt.Printf("%q\n", nonempty.Nonempty3(strs))
	fmt.Printf("%q\n", strs)
}

func TestNonempty4(t *testing.T) {
	fmt.Printf("%q\n", nonempty.Nonempty4(strs))
	fmt.Printf("%q\n", strs)
}

//var strs = []string{"", "", "", "", "", "", "", "A man", "A man", "", "", "", "", "", "", "", "", "", "", "A man", "A man", "", "", "", "", "", "", "", "", "", "", "A man", "A man", "", "", "", "", "", "", "", "", "", "", "A man", "A man", "", "", "", "", "", "", "", "", "", "", "A man", "A man", "", "", "", "", "", "", "", "", "", "", "A man", "A man", "", "", "", "", "", "", "", "", "", "", "A man", "A man", "", "", "", ""}
var strs = []string{
	"A man", "1", "", "", "", "", "", "", "", "", "", "A man", "A man", "2", "", "", "", "",
	"A man", "1", "", "", "", "", "", "", "", "", "", "A man", "A man", "2", "", "", "", "",
	"", "", "", "", "", "A man", "A man", "3", "A man", "A man", "3", "A man", "A man", "3",
	"A man", "A man", "3", "A man", "A man", "3", "A man", "A man", "3", "A man", "A man", "3",
	"A man", "A man", "3", "A man", "A man", "3", "", "", "", "", "", "", "", "", "", "A man",
	"A man", "4", "", "", "", "", "", "", "", "", "", "A man", "A man", "5", "", "", "", "",
	"", "", "", "", "", "A man", "A man", "6", "", "", "", "", "", "", "", "", "", "A man",
	"A man", "7", "", "", ""}

//!+bench
func BenchmarkNonempty0(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		nonempty.Nonempty0(strs)
	}
}

//!+bench
func BenchmarkNonempty1(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		nonempty.Nonempty1(strs)
	}
}

//!+bench
func BenchmarkNonempty2(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		nonempty.Nonempty2(strs)
	}
}

//!+bench
func BenchmarkNonempty3(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		nonempty.Nonempty3(strs)
	}
}

//!+bench
func BenchmarkNonempty4(b *testing.B) {
	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		nonempty.Nonempty4(strs)
	}
}

func BenchmarkArray(b *testing.B) {
	fmt.Println(b.N)
	arr := new([1 << 5]int)
	//for i:= range
	cnt := 0
	for i := 0; i < b.N; i++ {
		cnt += len(arr)
	}
	fmt.Println(cnt)
}

func BenchmarkSlice(b *testing.B) {
	fmt.Println(b.N)
	cnt := 0
	slc := make([]int, 1<<5)
	for i := 0; i < b.N; i++ {
		cnt += len(slc)
	}
	fmt.Println(cnt)
}
