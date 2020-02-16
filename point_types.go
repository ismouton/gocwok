package main

import (
	"fmt"
)

// Bounds represents the SW and NE bounds of a shape
type Bounds [2]*Point

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

// IsEqual determines if two points are equal
func (p *Point) IsEqual(c *Point) bool {
	if c == nil && p == nil {
		return true
	}

	if c == nil || p == nil {
		return false
	}

	return FloatEquality(p.X, c.X, .000001) && FloatEquality(p.Y, c.Y, .000001)
}
