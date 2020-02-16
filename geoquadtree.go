package main

// GeoQuadTree used for indexing GeoShapes
type GeoQuadTree struct {
	FeatureCount int64
	Root         *GeoQuadTreeNode
}

// GeoQuadTreeNode used for holding data
type GeoQuadTreeNode struct {
	Alpha   *GeoQuadTreeNode // top right
	Beta    *GeoQuadTreeNode // top left
	Gamma   *GeoQuadTreeNode // bottom left
	Delta   *GeoQuadTreeNode // bottom right
	Featues []Feature
	Bounds  Bounds
}

// Add adds a Feature to a GeoQuadTree
func (q *GeoQuadTree) Add(f *Feature) *GeoQuadTreeNode {
	// TODO write this
	return q.Root
}

// CheckBounds determines if `a` is a child of `q`
func (q *GeoQuadTreeNode) CheckBounds(f *Feature) bool {
	fSW := f.Bounds[0]
	fNE := f.Bounds[1]
	qSW := q.Bounds[0]
	qNE := q.Bounds[1]

	SWInBounds := fSW.X >= qSW.X && fSW.Y >= qSW.Y && fSW.X <= qNE.X && fSW.Y <= qNE.Y
	NEInBounds := fNE.X <= qNE.X && fNE.Y <= qNE.Y && fNE.X >= qSW.X && fNE.Y >= qSW.Y

	return SWInBounds && NEInBounds
}
