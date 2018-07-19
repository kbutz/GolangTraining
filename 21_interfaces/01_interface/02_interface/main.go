package main

import "fmt"

type square struct {
	side float64
}

type rectangle struct {
	width  float64
	length float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func(z rectangle) area() float64 {
	return z.width * z.length
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s)
	info(s)

	r := rectangle{width: 10, length: 20}
	fmt.Printf("%T\n", r)
	info(r)
}
