package main

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	oracleTest := func(
		oracle *Point,
		c0 LineSegment,
		c1 LineSegment,
	) {
		result := FindIntersection(c0, c1)

		// test for nils
		if oracle == nil && result == nil {
			return
		} else if oracle == nil && result != nil {
			t.Errorf("Expected no intersection but found: %s\n", result)
			return
		} else if oracle != nil && result == nil {
			t.Errorf("Expected an intersection of (%s) but found none.\n", oracle)
			return
		}

		theshold := 0.000001
		if !FloatEquality(result.X, oracle.X, theshold) || !FloatEquality(result.Y, oracle.Y, theshold) {
			t.Errorf("Got: %s; Expected: %s;\n", result, oracle)
		}
	}

	// intersect
	oracleTest(
		&Point{0.5, 0.5},
		[2]*Point{
			&Point{0, 0},
			&Point{1, 1},
		},
		[2]*Point{
			&Point{1, 0},
			&Point{0, 1},
		},
	)

	// share one vertex -- this should not count as an intersection
	oracleTest(
		nil,
		[2]*Point{
			&Point{0, 0},
			&Point{1, 1},
		},
		[2]*Point{
			&Point{1, 1},
			&Point{2, 2},
		},
	)

	// test with an undefined line
	oracleTest(
		&Point{X: 0.500000, Y: 1.250000},
		[2]*Point{
			&Point{0.5, 0},
			&Point{0.5, 5},
		},
		[2]*Point{
			&Point{0, 1},
			&Point{2, 2},
		},
	)

	// test with two undefined lines where one line is a segment of the other
	oracleTest(
		nil,
		[2]*Point{
			&Point{0.5, 0},
			&Point{0.5, 5},
		},
		[2]*Point{
			&Point{0.5, 1},
			&Point{0.5, 4},
		},
	)

	// test with two parallel lines
	oracleTest(
		nil,
		[2]*Point{
			&Point{0, 0},
			&Point{5, 3},
		},
		[2]*Point{
			&Point{1, 0},
			&Point{6, 3},
		},
	)
}
