package main

import (
	"testing"
)

func TestGeoQuadTreeBounds(t *testing.T) {
	oracleTest := func (
		oracle bool,
		c0 Bounds,
		c1 Bounds,
	) {
		q := &GeoQuadTreeNode{Bounds: c0}
		f := &Feature{Bounds: c1}
		fits := q.CheckBounds(f)

		if fits != oracle {
			t.Errorf("Expected %t; Got: %t", oracle, fits)
		}
	}

	oracleTest(
		true,
		[2]*Point{
			&Point{0,0},
			&Point{1,1},
		},
		[2]*Point{
			&Point{0.1,0.1},
			&Point{0.9,0.9},
		},
	)

	oracleTest(
		false,
		[2]*Point{
			&Point{0,0},
			&Point{1,1},
		},
		[2]*Point{
			&Point{-0.1,0.1},
			&Point{0.9,0.9},
		},
	)

	oracleTest(
		false,
		[2]*Point{
			&Point{0,0},
			&Point{1,1},
		},
		[2]*Point{
			&Point{0.1,-0.1},
			&Point{0.9,0.9},
		},
	)

	oracleTest(
		false,
		[2]*Point{
			&Point{0,0},
			&Point{1,1},
		},
		[2]*Point{
			&Point{0.1,0.1},
			&Point{-0.9,0.9},
		},
	)

	oracleTest(
		false,
		[2]*Point{
			&Point{0,0},
			&Point{1,1},
		},
		[2]*Point{
			&Point{0.1,0.1},
			&Point{0.9,-0.9},
		},
	)
}