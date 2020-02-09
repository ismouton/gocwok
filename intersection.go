package main

import "math"

import "fmt"

// FindIntersection finds the point of intersection of two line segemnts
func FindIntersection(
	lineSegment0 [2]Point,
	lineSegment1 [2]Point,
) *Point {
	// line segment 0
	Ax := lineSegment0[0].X
	Ay := lineSegment0[0].Y
	Bx := lineSegment0[1].X
	By := lineSegment0[1].Y

	// line segment 1
	Cx := lineSegment1[0].X
	Cy := lineSegment1[0].Y
	Dx := lineSegment1[1].X
	Dy := lineSegment1[1].Y

	var distAB float64
	var theCos float64
	var theSin float64
	var newX float64
	var ABpos float64

	//  Fail if both lines are undefined.
	if Ax == Bx && Ay == By && Cx == Dx && Cy == Dy {
		fmt.Println("One line is undefined")
		return nil
	}

	// if either line is undefined then we have to use a different method
	// TODO write this case
	if Ax == Bx && Ay == By || Cx == Dx && Cy == Dy {
		// undefined line segment
		var uLineSegment *Point

		// defined line segment
		var dLineSegment *Point

		if Ax == Bx && Ay == By {
			uLineSegment = []*Point{
				&Point{X: Ax, Y: Ay},
				&Point{X: Bx, Y: By},
			}

			dLineSegment = []*Point{
				&Point{X: Cx, Y: Cy},
				&Point{X: Dx, Y: Dy},
			}
		} else { 
			uLineSegment = []*Point{
				&Point{X: Cx, Y: Cy},
				&Point{X: Dx, Y: Dy},
			}

			dLineSegment = []*Point{
				&Point{X: Ax, Y: Ay},
				&Point{X: Bx, Y: By},
			}
		}

		// determine if two line segments intersect
		if 

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

	//  Fail if the lines are parallel.
	if Cy == Dy {
		return nil
	}

	//  (3) Discover the position of the intersection point along line A-B.
	ABpos = Dx + (Cx-Dx)*Dy/(Dy-Cy)

	//  (4) Apply the discovered position to line A-B in the original coordinate system.
	X := Ax + ABpos*theCos
	Y := Ay + ABpos*theSin

	//  Success.
	return &Point{
		X,
		Y,
	}
}

// isPointContainedByShape determines if a point is fully enclosed by `testShape`
func isPointContainedByShape(testShape *GeoNode, point *Point) bool {

	return true
}
