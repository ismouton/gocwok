package main

import (
	"strings"
	"sync"

	"github.com/jonas-p/go-shp"
)

// BuildFeatures indexes a shape file into provided data structure
func BuildFeatures(filename *string) ([]*Feature, error) {
	features := make([]*Feature, 0)
	shape, err := shp.OpenZip(*filename)
	featureChan := make(chan *Feature)
	done := make(chan bool)
	var wg sync.WaitGroup

	if err != nil {
		return nil, err
	}

	go func() {
		for feature := range featureChan {
			features = append(features, feature)
		}

		done <- true
	}()

	for shape.Next() {
		wg.Add(1)
		_, cur := shape.Shape()
		go func(cur shp.Shape) {
			defer wg.Done()

			fields := make(map[string]string)
			for k, f := range shape.Fields() {
				val := shape.Attribute(k)
				fields[strings.ToLower(f.String())] = val
			}

			bounds := extractBBox(cur)
			shapes := extractShapes(cur)
			centroid, area, err := FindCentroid(shapes)

			l := len(shapes)

			feature := &Feature{
				GeoShapeCount: int64(l),
				GeoShapes:     shapes,
				Bounds:        bounds,
				Properties:    fields,
			}

			featureChan <- feature
		}(cur)
	}

	wg.Wait()
	close(featureChan)

	<-done
	return features, nil
}
