// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var m = flag.Int("m", 0, "test")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	a := *sep
	b := *n
	fmt.Print(strings.Join(flag.Args(), a))
	if !b && *m > 0 {
		fmt.Println()
	}
}

//!-
