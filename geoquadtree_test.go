package main

import (
	"testing"
)

func TestGeoQuadTreeBounds(t *testing.T) {
	oracleTest := func (
		oracle bool,
		bounds0 Bounds,
		bounds1 Bounds,
	) {
		q := &GeoQuadTreeNode{Bounds: bounds0}
		f := &Feature{Bounds: bounds1}
		fits := q.CheckBounds(f)

		if fits != oracle {
			t.Errorf("Expected %t; Got: %t", oracle, fits)
		}
	}

	oracleTest(
		true,
		[]Point{
			Point{0,0},
			Point{1,1},
		},
		[]Point{
			Point{0.1,0.1},
			Point{0.9,0.9},
		},
	)

	oracleTest(
		false,
		[]Point{
			Point{0,0},
			Point{1,1},
		},
		[]Point{
			Point{-0.1,0.1},
			Point{0.9,0.9},
		},
	)

	oracleTest(
		false,
		[]Point{
			Point{0,0},
			Point{1,1},
		},
		[]Point{
			Point{0.1,-0.1},
			Point{0.9,0.9},
		},
	)

	oracleTest(
		false,
		[]Point{
			Point{0,0},
			Point{1,1},
		},
		[]Point{
			Point{0.1,0.1},
			Point{-0.9,0.9},
		},
	)

	oracleTest(
		false,
		[]Point{
			Point{0,0},
			Point{1,1},
		},
		[]Point{
			Point{0.1,0.1},
			Point{0.9,-0.9},
		},
	)
}