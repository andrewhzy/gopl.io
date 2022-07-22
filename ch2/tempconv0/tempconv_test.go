// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package tempconv

import (
	"fmt"
	"math"
	"testing"
)

func Example_one() {
	{
		//!+arith
		fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	}
	// Output:
	// 100
	// 180
}

func Example_two() {
	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c) // "100°C"
	fmt.Println(c)        // "100°C"
	fmt.Printf("%g\n", c) // "100"; does not call String
	fmt.Println(float64(c))
	fmt.Println(Fahrenheit(Celsius(9)))
	//Output:
	//100°C
	//100°C
	//100°C
	//100°C
	//100
	//100
}

func TestCToF(t *testing.T) {
	fmt.Println(float64(Celsius(Fahrenheit(float64(1<<63 - 1)))))
	fmt.Println(int(float64(1 << 62)))
	fmt.Println(1 << 62)
	fmt.Println(int(float64(math.MaxInt)))
	fmt.Println(1<<62 - 10)
	fmt.Println((float64(1<<63 - 10)))
	fmt.Println(float64(1<<69 - 10))
}
