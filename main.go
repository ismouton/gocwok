package main

import "fmt"

func main() {
	node0 := &GeoNode{Coordinates: &Point{0, 0}}
	node1 := &GeoNode{Coordinates: &Point{1, 1}}
	node2 := &GeoNode{Coordinates: &Point{2, 2}}
	node3 := &GeoNode{Coordinates: &Point{3, 3}}

	node0.InsertAfter(node1).InsertAfter(node2).InsertAfter(node3)

	ringClone := node0.CloneRing()

	node1.Unlink()

	node3.InsertAfter(&GeoNode{Coordinates: &Point{2.5, 2.5}})
	node2.InsertBefore(&GeoNode{Coordinates: &Point{1.5, 1.5}})

	fmt.Println("Node 0:")
	node0.Print()

	fmt.Println("Cloned Node:")
	ringClone.Print()

	fmt.Println("Concat Cloned Node:")
	// ringClone.InsertAfter(node0.CloneRing().BreakRingBefore())
	ringClone.BreakRingBefore().Print()

	fmt.Println("Node 0:")
	node0.Print()
}
