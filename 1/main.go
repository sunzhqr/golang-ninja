package main

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Subtract() int {
	return n.num1 - n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Division() float32 {
	if n.num2 != 0 {
		return float32(n.num1 / n.num2)
	}
	return 0
}

type NumbersInterface interface {
	Sum() int
	Subtract() int
	Multiply() int
	Division() float32
}

func main() {
	var i NumbersInterface
	nums1 := Numbers{10, 0}
	//nums2 := Numbers{10, 0}
	i = nums1
	fmt.Println("Sum", i.Sum())
	fmt.Println("Subtract", i.Subtract())
	fmt.Println("Multiply", i.Multiply())
	fmt.Println("Division", i.Division())
}
