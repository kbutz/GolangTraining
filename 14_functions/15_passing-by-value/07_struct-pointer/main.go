package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Todd", 44}
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x8201e4120

	changeString(&c1.name)
	fmt.Println(c1)
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Todd 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x8201e4120

}

func changeString(n *string) {
	fmt.Println("---- change string ----")
	fmt.Println(n)
	fmt.Println(*n)
	*n = "Name Change"


}
