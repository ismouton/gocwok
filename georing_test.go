package main

import "testing"

func TestGeoRing(t *testing.T) {
	oracleTest := func(oracle []Point, ring *GeoNode) {
		cur := ring
		first := cur

		i := 0
		for {
			if cur.Coordinates.X != oracle[i].X || cur.Coordinates.Y != oracle[i].Y {
				t.Errorf(
					"X,Y == %f, %f; want %f, %f",
					cur.Coordinates.X,
					cur.Coordinates.Y,
					oracle[i].X,
					oracle[i].Y,
				)
			}

			cur = cur.Next
			i++
			if cur == first {
				break
			}
		}
	}

	node0 := &GeoNode{Coordinates: &Point{0, 0}}
	node1 := &GeoNode{Coordinates: &Point{1, 1}}
	node2 := &GeoNode{Coordinates: &Point{2, 2}}
	node3 := &GeoNode{Coordinates: &Point{3, 3}}

	node0.InsertAfter(node1).InsertAfter(node2).InsertAfter(node3)
	node0.Print()
	ringClone := node0.CloneRing()

	l := node0.Len()
	if l != 4 {
		t.Errorf("Len == %d; want 4", l)
	}

	node2.Unlink()

	l = node0.Len()
	if l != 3 {
		t.Errorf("Len == %d; want 3", l)
	}

	node0.InsertBefore(&GeoNode{Coordinates: &Point{-1, -1}})

	// test ring to ensure data was inserted in proper order
	oracleTest(
		[]Point{
			Point{0, 0},
			Point{1, 1},
			Point{3, 3},
			Point{-1, -1},
		},
		node0,
	)

	// lets ensure the clone was not mutated by operations on other ring
	oracleTest(
		[]Point{
			Point{0, 0},
			Point{1, 1},
			Point{2, 2},
			Point{3, 3},
		},
		ringClone,
	)
}
