package main

import (
	"fmt"
	"testing"
)

func Test0(t *testing.T) {
	ints := []int{1, 2}
	fmt.Println(cap(ints))
	ints = append(ints, 3, 4, 5)
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
}
