package main

import (
	"fmt"

	"github.com/jonas-p/go-shp"
)

func main() {
	states := "shp/states.zip"
	shape, err := shp.OpenZip(states)

	if err != nil {
		panic(err)
	}

	for shape.Next() {
		fields := extractFields(shape)
		bounds := extractBBox(shape)
		shapes := extractShapes(shape)
		l := len(shapes)
		fmt.Printf("Found %d shapes in %s!\n", l, fields["name"])

		feature := &Feature{
			GeoShapeCount: int64(l),
			GeoShapes:     shapes,
			Bounds:        bounds,
			Properties:    fields,
		}

		if feature == nil {

		}
	}

	return

	node0 := &GeoNode{Coordinates: &Point{0, 0}}
	node1 := &GeoNode{Coordinates: &Point{1, 1}}
	node2 := &GeoNode{Coordinates: &Point{2, 2}}
	node3 := &GeoNode{Coordinates: &Point{3, 3}}
	node0.InsertAfter(node1).InsertAfter(node2).InsertAfter(node3)
	ringClone := node0.CloneRing()

	// node1.Unlink()

	node3.InsertBefore(&GeoNode{Coordinates: &Point{2.5, 2.5}})
	node1.InsertAfter(&GeoNode{Coordinates: &Point{1.5, 1.5}})

	fmt.Println("Node 0:")
	node0.Print()

	fmt.Println("Cloned Node:")
	ringClone.Print()

	fmt.Println("Concat Cloned Node:")
	ringClone.InsertBefore(node3.CloneRing().BreakRingBefore())
	ringClone.Print()

	fmt.Println("Node 0:")
	node0.Print()

	intersection := FindIntersection(
		[2]*Point{
			&Point{0, 0},
			&Point{1, 1},
		},
		[2]*Point{
			&Point{1, 0},
			&Point{0, 1},
		},
	)

	if intersection != nil {
		fmt.Printf("Intersection: %f, %f", intersection.X, intersection.Y)
	}

}
