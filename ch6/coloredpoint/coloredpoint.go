package main

import (
	"fmt"
	"image/color"
	"math"
)

// Point Basic test struct
type Point struct {
	X, Y float64
}

// Distance but it is a standalone function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy scales the point by modifying the underlying struct
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Distance but its attached to the point method allowing for dot chains
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path is a journey connecting points together with straight lines
type Path []Point

// Distance , a method that returns the distance of a path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		// make sure we start at the second item and count back
		if i > 0 {
			// take current point and previous point and measure distance
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// ColoredPoint a colored point in space
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
	p := Point{0, 0}
	// fmt.Println(p.Distance(cp)) // Problem because of a "has-a" relationship
	fmt.Println(cp.Distance(p)) // works because we added a wrapper func

	q := Point{10, 10}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
}
