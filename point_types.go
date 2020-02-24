package main

import (
	"fmt"
)

// Bounds represents the SW and NE bounds of a shape
type Bounds [2]*Point

// Triad represents 3 points
type Triad [3]*Point

// LineSegment represents a simple line segement
type LineSegment [2]*Point

// Point represents a point in 2d space
type Point struct {
	X float64
	Y float64
}

// String transforms a point into a string
func (p *Point) String() string {
	return fmt.Sprintf("X:%f, Y:%f", p.X, p.Y)
}
