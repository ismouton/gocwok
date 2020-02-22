package main

import (
	"github.com/jonas-p/go-shp"
)

func extractBBox(p shp.Shape) Bounds {
	polygon := p.(*shp.Polygon)

	bbox := polygon.BBox()

	bounds := Bounds{
		&Point{
			bbox.MaxY,
			bbox.MaxX,
		},
		&Point{
			bbox.MinY,
			bbox.MinX,
		},
	}

	return bounds
}

func extractShapes(p shp.Shape) []*GeoNode {
	geoShapes := make([]*GeoNode, 0)
	polygon := p.(*shp.Polygon)

	shapeCount := 0
	var first *Point
	var currentGeoNode *GeoNode
	for _, s := range polygon.Points {
		current := &Point{s.X, s.Y}

		if current.IsEqual(first) {
			shapeCount++
			first = nil
			geoShapes = append(geoShapes, currentGeoNode)

			continue
		}

		if first == nil {
			first = current
			currentGeoNode = &GeoNode{Coordinates: first}
		} else {
			currentGeoNode = currentGeoNode.InsertAfter(&GeoNode{Coordinates: current})
		}
	}

	return geoShapes
}
