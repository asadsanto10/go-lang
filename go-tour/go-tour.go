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

// switch with no condition
func switchCondition() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("good afternoon")
	case time.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")
	}
}

func defers() {
	// defer fmt.Println("with defers 1 ")
	// defer fmt.Println("with defers 2 ")
	// defer fmt.Println("with defers 3 ")
	// defer fmt.Println("with defers 4 ")
	// defer fmt.Println("with defers 5 ")
	// fmt.Println("without defers 0")

	fmt.Println("contining")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}

// Pointers
func pointers() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(p)
	fmt.Println(i)

	p = &j
	*p = *p / 37

	fmt.Println(j)
}

// Structs
func structsType() {
	type Vertex struct {
		x int
		y int
	}

	v := Vertex{1, 3}
	v.x = 4

	fmt.Println(v)

	// pointer to struct
	p := &v
	ab := &p
	p.x = 585
	fmt.Println("change value")
	fmt.Println(p)
	fmt.Println(v)
	fmt.Println("ab", ab)

	// Struct Literals

	type Column struct {
		a, v int
	}

	var (
		v1  = Column{10, 11}
		v2  = Column{a: 15}
		v3  = Column{}
		poo = &v1
	)

	// fmt.Println(v1, poo, v2, V3)
	fmt.Println(v1, v2, v3, *poo)

}

// arrays
func arrays() {
	var a [2]string
	fmt.Println("::::::::Arrays::::::::")
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 7, 11, 13}
	
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

	// fmt.Println(time.Now().ISOWeek())

	switchCondition()

	// defers()
	pointers()

	structsType()
	arrays()

}
