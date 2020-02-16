package main

import (
	"math"
)

// FindIntersection finds the point of intersection of two line segemnts
func FindIntersection(
	l0 LineSegment,
	l1 LineSegment,
) *Point {
	// line segment 0
	Ax := l0[0].X
	Ay := l0[0].Y
	Bx := l0[1].X
	By := l0[1].Y

	// line segment 1
	Cx := l1[0].X
	Cy := l1[0].Y
	Dx := l1[1].X
	Dy := l1[1].Y

	var distAB float64
	var theCos float64
	var theSin float64
	var newX float64
	var ABpos float64

	//  Fail if either line segment is zero-length.
	if Ax == Bx && Ay == By || Cx == Dx && Cy == Dy {
		return nil
	}

	//  Fail if the segments share an end-point.
	if Ax == Cx && Ay == Cy || Bx == Cx && By == Cy ||
		Ax == Dx && Ay == Dy || Bx == Dx && By == Dy {
		return nil
	}

	//  (1) Translate the system so that point A is on the origin.
	Bx -= Ax
	By -= Ay
	Cx -= Ax
	Cy -= Ay
	Dx -= Ax
	Dy -= Ay

	//  Discover the length of segment A-B.
	distAB = math.Sqrt(Bx*Bx + By*By)

	//  (2) Rotate the system so that point B is on the positive X axis.
	theCos = Bx / distAB
	theSin = By / distAB
	newX = Cx*theCos + Cy*theSin
	Cy = Cy*theCos - Cx*theSin
	Cx = newX
	newX = Dx*theCos + Dy*theSin
	Dy = Dy*theCos - Dx*theSin
	Dx = newX

	//  Fail if segment C-D doesn't cross line A-B.
	if Cy < 0. && Dy < 0. || Cy >= 0. && Dy >= 0. {
		return nil
	}

	//  (3) Discover the position of the intersection point along line A-B.
	ABpos = Dx + (Cx-Dx)*Dy/(Dy-Cy)

	//  Fail if segment C-D crosses line A-B outside of segment A-B.
	if ABpos < 0. || ABpos > distAB {
		return nil
	}

	//  (4) Apply the discovered position to line A-B in the original coordinate system.
	X := Ax + ABpos*theCos
	Y := Ay + ABpos*theSin

	//  Success.
	return &Point{
		X: X,
		Y: Y,
	}
}

// IsPointContainedByShape determines if point is fully enclosed by testShape
func IsPointContainedByShape(testShape *GeoNode, point *Point) bool {
	max := math.Pow(2, 32)
	ray := [2]*Point{}
	ray[0] = point
	ray[1] = &Point{
		X: point.X,
		Y: max,
	}

	intersectionCount := 0

	first := testShape
	cur := first

	// test all line segments against ray
	for {
		currentLinesegment := [2]*Point{
			cur.Coordinates,
			cur.Next.Coordinates,
		}

		if FindIntersection(ray, currentLinesegment) != nil {
			intersectionCount++
		}

		cur = cur.Next
		if cur == first {
			break
		}
	}

	// even number of intersections means outside shape
	if intersectionCount%2 == 0 {
		return false
	}

	return true
}
