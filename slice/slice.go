package main

import "fmt"

func main() {
	// slice
	mySlice := [5]int{1, 2, 3, 4}
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
