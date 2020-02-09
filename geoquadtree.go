package main

// GeoQuadTree used for indexing GeoShapes
type GeoQuadTree struct {
	ShapeCount int64
	Root       *GeoNode
}

type GeoQuadTreeNode struct {
	Alpha   *GeoQuadTreeNode // top right
	Beta    *GeoQuadTreeNode // top left
	Gamma   *GeoQuadTreeNode // bottom left
	Delta   *GeoQuadTreeNode // bottom right
	Featues []Feature
}
