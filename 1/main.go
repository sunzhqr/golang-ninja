package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	//square := Square{5}
	circle := Circle{5}
	//printArea(square)
	//printArea(circle)
	printInterface(Circle{5}) //unknown
	var s struct{}
	printInterface(s) //struct
	printInterface("Sanzhar")
	printInterface(circle)
}

func printArea(shape Shape) {
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case struct{}:
		fmt.Println("struct")
	case string:
		fmt.Println("string")
	case float32:
		fmt.Println("float")
	default:
		fmt.Println("unknown")
	}
}
