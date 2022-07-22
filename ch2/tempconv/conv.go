// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv111

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts a Fahrenheit temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

//!-

// CToF converts a Celsius temperature to Fahrenheit.
func (c Celsius) ToF() Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func (f Fahrenheit) ToC() Celsius { return Celsius((f - 32) * 5 / 9) }

//!-
