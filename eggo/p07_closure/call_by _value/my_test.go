package call_by__value

import (
	"fmt"
	"testing"
)

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "Hi! " + a.name
	return a.name
}

func Test0(t *testing.T) {
	a := A{name: "eggo"}
	fmt.Println(a.Name())
	fmt.Println(A.Name(a))
}
