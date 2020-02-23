package main

import (
	"errors"
	"math"

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

// FindAreaOfTriad returns the area of a triad
func FindAreaOfTriad(t Triad) float64 {
	a := FindDistance(t[0], t[1])
	b := FindDistance(t[1], t[2])
	c := FindDistance(t[0], t[2])
	s := (a + b + c) / 2

	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// FindArea returns the area of a geonode
func FindArea(g *GeoNode) (float64, error) {
	// radius of earth in meters
	const r = 6378137

	convertToRadians := func(input float64) float64 {
		return input * math.Pi / 180
	}

	cur := g
	first := g
	area := 0.

	i := 0
	for {
		p1 := cur.Coordinates
		p2 := cur.Next.Coordinates

		area += convertToRadians(p2.X-p1.X) * (2 + math.Sin(convertToRadians(p1.Y)) + math.Sin(convertToRadians(p2.Y)))

		cur = cur.Next
		if cur == first {
			break
		}
		i++
	}

	if i < 2 {
		return 0, errors.New("shape has less than 3 points")
	}

	area *= r * r / 2

	return math.Abs(area), nil
}

// FindDistance finds distance between two points and returns value
func FindDistance(p1 *Point, p2 *Point) float64 {
	lat1 := p1.X
	lat2 := p2.X
	lng1 := p1.Y
	lng2 := p2.Y

	radlat1 := math.Pi * lat1 / 180
	radlat2 := math.Pi * lat2 / 180

	theta := lng1 - lng2
	radtheta := math.Pi * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	return dist
}

// FindCentroid finds the centroid
func FindCentroid(g *GeoNode) (*Point, error) {
	// TODO Write this
	return &Point{0, 0}, nil
}

// FindCentroidAndAreaOfSet finds the centroid of shape with most area in a set of shapes
func FindCentroidAndAreaOfSet(s []*GeoNode) (float64, *Point, error) {
	maxArea := 0.
	var maxShape *GeoNode
	for _, shape := range s {
		area, err := FindArea(shape)

		if err != nil {
			return 0, nil, err
		}

		if area > maxArea {
			maxShape = shape
			maxArea = area
		}
	}

	centroid, err := FindCentroid(maxShape)

	if err != nil {
		return 0, nil, err
	}

	return maxArea, centroid, nil
}
