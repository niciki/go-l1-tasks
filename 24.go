package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены
в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) GetDistance(p2 Point) float64 {
	// считаем расстояние по формуле из курса аналитической геометрии
	return math.Sqrt(float64((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y)))
}

func main() {
	var x1, x2, y1, y2 int
	fmt.Print("Please input coordinates of Point: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Please input coordinates of Point: ")
	fmt.Scan(&x2, &y2)
	point1 := NewPoint(x1, y1)
	point2 := NewPoint(x2, y2)
	fmt.Println("Distance between points: ", point1.GetDistance(point2))
}
