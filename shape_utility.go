package main

import (
	"strings"

	"github.com/jonas-p/go-shp"
)

func extractFields(shape *shp.ZipReader) map[string]string {
	fields := shape.Fields()

	tempFieldsStruct := make(map[string]string)
	for k, f := range fields {
		val := shape.Attribute(k)
		tempFieldsStruct[strings.ToLower(f.String())] = val
	}

	return tempFieldsStruct
}

func extractBBox(shape *shp.ZipReader) Bounds {
	_, p := shape.Shape()
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

func extractShapes(shape *shp.ZipReader) []*GeoNode {
	geoShapes := make([]*GeoNode, 0)
	_, p := shape.Shape()
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
