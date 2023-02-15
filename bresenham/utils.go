package main

import "strconv"

// Abs returns the absolute value of an int
// Rough implementation of math.Abs()
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// StoFtoI reads a string as a float64, and returns an int
func StoFtoI(s string) int {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return int(f)
}
