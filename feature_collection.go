package main

import "github.com/jonas-p/go-shp"

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
func (f *FeatureCollection) SaveToShapeFile(filename string) error {
	// create and open a shapefile for writing points
	shape, err := shp.Create(filename, shp.POLYGON)
	if err != nil {
		return err
	}
	defer shape.Close()

	return nil
}
