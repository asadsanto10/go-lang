package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

// multiple return type
func swap(x int, y string) (string, int) {
	return y, x
}

// Named return values
func testname(x string) (a string) {
	a = x
	return
}

func split(sum int) (b, c int) {
	b = sum * 2
	c = sum * 3
	return
}

// type conversions
func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
}

func main() {
	fmt.Println("random number is", rand.Intn(10))
	fmt.Println("math number is", math.Sqrt(7))
	fmt.Println("add is::", add(2, 2))

	a, b := swap(14, "asad")
	fmt.Println(a, b)

	fmt.Println(testname("test name returns value"))

	fmt.Println(split(10))

	typeConversion()

}
