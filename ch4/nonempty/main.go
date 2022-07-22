// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package nonempty

import (
	"fmt"
)

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func Nonempty0(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func Nonempty1(strings []string) []string {
	i := 0
	for i = range strings {
		if len(strings[i]) > 0 {
			break
		}
	}
	strings = strings[i:]
	for i = range strings {
		if len(strings[i]) == 0 {
			break
		}
	}
	j := i + 1
	for j < len(strings) {
		if len(strings[j]) == 0 {
			j++
			continue
		}
		strings[i] = strings[j]
		i++
		j++
	}
	return strings[:i]
}

//!-nonempty

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", Nonempty0(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)            // `["one" "three" "three"]`
	//!-main
}

func Nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func Nonempty3(strings []string) []string {
	i := 0
	for i = range strings {
		if len(strings[i]) > 0 {
			break
		}
	}
	strings = strings[i:]
	for i = range strings {
		if len(strings[i]) == 0 {
			break
		}
	}
	out := strings[:i]
	strings = strings[i:]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func Nonempty4(strings []string) []string {
	i := 0
	for i = range strings {
		if len(strings[i]) > 0 {
			break
		}
	}
	res := strings[i:]
	for i = range strings {
		if len(strings[i]) == 0 {
			break
		}
	}
	end := i
	strings = res[i+1:]
	i, j := 0, 0
	for i < len(strings) {
		if len(strings[i]) == 0 {
			i++
			continue
		}
		j = i + 1
		for j < len(strings) {
			if len(strings[j]) == 0 {
				break
			}
			j++
		}
		n := copy(res[end:], strings[i:j])
		end += n
		i = j + 1
	}
	return res[:end]
}
