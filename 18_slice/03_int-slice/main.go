package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 10)

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	capacity := cap(mySlice)

	for i := 0; i < 2000000; i++ {
		mySlice = append(mySlice, i)
		if cap(mySlice) > capacity {
			capacity = cap(mySlice)
			fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
		}
	}
}
