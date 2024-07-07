package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
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

// for loop
func forLoop() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("For loop::", sum)
}

// For continued
func forCounted() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("forCounted:: ", sum)
}

// For-each range loop
func forRange() {
	texts := []string{"asad", "santo"}
	for i, s := range texts {
		fmt.Println(i, s)
	}
}

// exit loop
func exitLoop() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("exit loop", sum)
}

// if statement
func ifStatement() {
	if 7%2 == 0 {
		fmt.Println("7 is event")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println(sqrt(2), sqrt(-4))

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}

// If with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Switch
func switchStatement() {
	fmt.Println("GO lang runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")

	case "linux":
		fmt.Println("Linux.")
	default:

		fmt.Printf("%s.\n", os)
	}

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

	forLoop()
	forCounted()
	forRange()
	exitLoop()
	ifStatement()
	fmt.Println(pow(2, 2, 10))
	switchStatement()

	fmt.Print(time.Now().ISOWeek())

}
