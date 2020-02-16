package main

import (
	"testing"
)

func TestFindIntersection(t *testing.T) {
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

	// test where line intersects but segment does not
	oracleTest(
		nil,
		[2]*Point{
			&Point{0, 0},
			&Point{5, 5},
		},
		[2]*Point{
			&Point{0, 30},
			&Point{30, 0},
		},
	)

	// test case with 0 length line
	oracleTest(
		nil,
		[2]*Point{
			&Point{0, 0},
			&Point{5, 5},
		},
		[2]*Point{
			&Point{1, 1},
			&Point{1, 1},
		},
	)
}

func TestIsPointContainedByShape(t *testing.T) {
	oracleTest := func(
		oracle bool,
		g *GeoNode,
		p *Point,
	) {
		result := IsPointContainedByShape(g, p)
		if result != oracle {
			t.Errorf("Expected %t; Got %t;\n", oracle, result)
		}
	}

	node0 := &GeoNode{Coordinates: &Point{0, 0}}
	node1 := &GeoNode{Coordinates: &Point{3, 0}}
	node2 := &GeoNode{Coordinates: &Point{3, 3}}
	node3 := &GeoNode{Coordinates: &Point{0, 3}}
	node0.InsertAfter(node1).InsertAfter(node2).InsertAfter(node3)

	// test case for point in shape
	oracleTest(true, node0, &Point{1.5, 1.5})

	// test case for point NOT in shape
	oracleTest(false, node0, &Point{5, 5})

	// shared vertex should not count as in shape
	oracleTest(false, node0, &Point{3, 3})
}
