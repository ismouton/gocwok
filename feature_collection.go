package main

import (
	"fmt"

	"github.com/jonas-p/go-shp"
)

// Feature represents a complete set of GeoShapes that represent a feature (ie County)
type Feature struct {
	GeoShapeCount int64
	GeoShapes     []*GeoNode
	Bounds        Bounds
	Properties    map[string]string
}

// Clone returns a deep clone of the feature
func (f *Feature) Clone() *Feature {
	// shallow copy feature
	clonedFeature := &Feature{}
	*clonedFeature = *f
	l := len(f.GeoShapes)
	clonedGeoShapes := make([]*GeoNode, l)

	for i, geoShape := range f.GeoShapes {
		clonedGeoShapes[i] = geoShape.Clone()
	}

	clonedFeature.GeoShapes = clonedGeoShapes

	return clonedFeature
}

// FeatureCollection is a collection of features that belong to a set
type FeatureCollection struct {
	Features []*Feature
}

// Clone returns a deep clone of the entire feature collection
func (f *FeatureCollection) Clone() *FeatureCollection {
	l := len(f.Features)
	clonedFeatures := make([]*Feature, l)

	for i, feature := range f.Features {
		clonedFeatures[i] = feature.Clone()
	}

	return &FeatureCollection{Features: clonedFeatures}
}

// SaveToShapeFile saves FeatureCollection to specified filename
func (f *FeatureCollection) SaveToShapeFile(filename *string) error {
	fmt.Printf("Writing file %s\n", *filename)
	// create and open a shapefile for writing points
	shape, err := shp.Create(*filename, shp.POLYGON)
	if err != nil {
		return err
	}
	defer shape.Close()

	for i0, feature := range f.Features {
		polyline := shp.Polygon{
			Points:   make([]shp.Point, 0),
			Parts:    []int32{0},
			NumParts: 1,
		}
		polyline.Parts = []int32{0}

		for partNumber, g := range feature.GeoShapes {
			cur := g
			first := g
			for {
				cur = cur.Next
				if *cur == *first {
					if partNumber < len(feature.GeoShapes)-1 {
						polyline.NumParts++
						polyline.Parts = append(polyline.Parts, int32(polyline.NumPoints))
					}

					break
				} else {
					polyline.NumPoints++
					polyline.Points = append(polyline.Points, shp.Point{X: cur.Previous.Coordinates.X, Y: cur.Previous.Coordinates.Y})
				}
			}
		}

		// fields to write
		fields := []shp.Field{
			// String attribute field with length 25
			shp.StringField("NAME", 25),
		}

		// setup fields for attributes
		shape.SetFields(fields)
		shape.Write(&polyline)
		shape.WriteAttribute(i0, 0, feature.Properties["name"])
	}

	return nil
}
