package main

import (
	"math"
)

func FloatEquality(
	f0 float64,
	f1 float64,
	threshold float64,
) bool {
	delta := math.Abs(f0 - f1)

	return delta < threshold
}
