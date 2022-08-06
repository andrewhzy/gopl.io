package call_by_value

import (
	"fmt"
	"testing"
)

type A struct {
	name string
}

func (a A) GetName() string {
	return a.name
}
func (pa *A) SetName() string {
	pa.name = "Hi! " + pa.name
	return pa.name
}

func Test0(t *testing.T) {
	a := A{name: "eggo"}
	pa := &a

	fmt.Println(pa.GetName())
	fmt.Println(a.SetName())
}
