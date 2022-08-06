package call_by_value

import (
	"fmt"
	"testing"
)

func B(a int) int {
	a++
	return a
}

func A(a int) {
	a++
	fmt.Println(a)
}

func Test0(t *testing.T) {
	a := 1
	defer A(B(a))
	a++
	fmt.Println(a)
}
