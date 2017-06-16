package main

import (
	"math"
)

func Sqrt(x float64) float64 {
	var x1 float64
	var x2 float64

	var f float64
	var dfdx float64

	if x < 0 {
		return math.NaN()
	}

	eps := 1e-5

	x1 = eps * 2
	x2 = 0

	for (x1-x2) > eps || (x2-x1) > eps {
		x2 = x1

		f = x2*x2 - x

		dfdx = 2 * x2

		x1 = x2 - f/dfdx
	}

	return x1
}

func main() {}
