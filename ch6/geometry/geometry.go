package main

import (
	"fmt"
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

func main() {
	point1 := Point{1, 2}
	point2 := Point{4, 6}
	fmt.Println(Distance(point1, point2))
	fmt.Println(point1.Distance(point2))
	fmt.Printf("Point1: %v\n", point1)
	point1.ScaleBy(2)
	fmt.Printf("Point1: %v\n", point1)
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
