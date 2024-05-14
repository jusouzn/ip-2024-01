package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, z float64
}

func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
}

func main() {
	var N int
	fmt.Scan(&N)

	points := make([]Point, N)
	for i := range points {
		fmt.Scan(&points[i].x, &points[i].y, &points[i].z)
	}

	for i := 0; i < N-1; i++ {
		d := distance(points[i], points[i+1])
		fmt.Printf("%.2f\n", d)
	}
}
