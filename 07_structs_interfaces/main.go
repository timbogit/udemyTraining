package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}


func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2.0
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * l + 2 * w
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	fmt.Println("end")
}
