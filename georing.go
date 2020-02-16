package main

import "fmt"

// Point represents a point in 2d space
type Point struct {
	X float64
	Y float64
}

// GeoNode repersents a single point in a GeoShape
type GeoNode struct {
	Previous    *GeoNode
	Next        *GeoNode
	Coordinates *Point
}

// Bounds represents the SW and NE bounds of a shape
type Bounds []Point

// Feature represents a complete set of GeoShapes that represent a feature (ie County)
type Feature struct {
	GeoShapeCount int64
	GeoShapes     []*GeoNode
	Bounds        Bounds
}

// Len returns the number of elements in the ring. Runs in O(n) time
func (g *GeoNode) Len() int {
	cur := g
	first := cur

	i := 0
	for {
		i++
		cur = cur.Next
		if cur == first {
			break
		}
	}

	return i
}

// Print prints all nodes
func (g *GeoNode) Print() {
	cur := g
	first := g

	for {
		fmt.Println(cur.Coordinates)

		cur = cur.Next
		if cur == first {
			break
		}
	}
}

// CloneNode returns a copy of GeoNode with Next and Previous set to nil
func (g *GeoNode) CloneNode() *GeoNode {
	n := &GeoNode{}
	*n = *g
	n.Previous = nil
	n.Next = nil

	return n
}

// BreakRingAfter breaks the ring afer this GeoNode creating a list returning first node after the break
func (g *GeoNode) BreakRingAfter() *GeoNode {
	beforeBreak := g
	afterBreak := g.Next

	// break the ring
	beforeBreak.Next = nil
	afterBreak.Previous = nil

	return afterBreak
}

// BreakRingBefore breaks the ring before this GeoNode creating a list returning first node after the break
func (g *GeoNode) BreakRingBefore() *GeoNode {
	beforeBreak := g.Previous
	afterBreak := g

	// break the ring
	beforeBreak.Next = nil
	afterBreak.Previous = nil

	return afterBreak
}

// CloneRing returns a copy of GeoNode
func (g *GeoNode) CloneRing() *GeoNode {
	clonedGeoNode := &GeoNode{}
	cur := &GeoNode{}

	*clonedGeoNode = *g

	// remove references to old ring
	clonedGeoNode.Next = nil
	clonedGeoNode.Previous = nil

	first := g
	*cur = *g.Next
	clonedCur := clonedGeoNode

	for {
		copiedCur := &GeoNode{}
		*copiedCur = *cur

		// break all links from node
		copiedCur.Next = nil
		copiedCur.Previous = nil

		clonedCur = clonedCur.InsertAfter(copiedCur)
		cur = cur.Next

		if cur == first {
			break
		}
	}

	return clonedGeoNode
}

// FindTerminiOfSegment returns the beginning and end of a segment
func (g *GeoNode) FindTerminiOfSegment() (*GeoNode, *GeoNode) {
	// lets find termini of `a`
	var beginning *GeoNode // beginning temrinus
	var end *GeoNode       // end terminus

	beginning = g
	for {
		if beginning == nil || beginning.Previous == nil {
			break
		}

		beginning = beginning.Previous
	}

	end = g
	for {
		if end == nil || end.Next == nil {
			break
		}

		end = end.Next
	}

	return beginning, end
}

// InsertBefore inserts a node or segment before this GeoNode
func (g *GeoNode) InsertBefore(a *GeoNode) *GeoNode {
	if g.Previous == nil {
		g.Previous = g
	}

	if g.Next == nil {
		g.Next = g
	}

	beginning, end := a.FindTerminiOfSegment()

	beginning.Previous = g.Previous

	g.Previous.Next = beginning
	g.Previous = end
	end.Next = g

	return a
}

// InsertAfter inserts a node or segment after this GeoNode
func (g *GeoNode) InsertAfter(a *GeoNode) *GeoNode {
	if g.Previous == nil {
		g.Previous = g
	}

	if g.Next == nil {
		g.Next = g
	}

	beginning, end := a.FindTerminiOfSegment()

	beginning.Previous = g
	end.Next = g.Next

	g.Next.Previous = end
	g.Next = beginning
	beginning.Previous = g

	return a
}

// Unlink removes node a
func (g *GeoNode) Unlink() {
	g.Next.Previous = g.Previous
	g.Previous.Next = g.Next

	// mark for gc
	g = nil
}
