package call_by_value

import (
	"testing"
)

type A struct {
	name string
}

func (a A) GetName() string {
	return a.name
}

func Name(a A) string {
	return a.name
}

func Test0(t *testing.T) {
	a := A{name: "eggo"}

	f1 := A.GetName
	f1(a)

	f2 := a.GetName
	f2()

	f3 := Name
	f3(a)
}
