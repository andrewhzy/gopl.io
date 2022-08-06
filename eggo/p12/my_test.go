package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type T int
type I interface {
	Name() string
}

func (i T) testT() {
}

func Test0(t *testing.T) {
	var rw io.ReadWriter
	r, ok := rw.(*os.File)
	fmt.Println(r, ok)
}
