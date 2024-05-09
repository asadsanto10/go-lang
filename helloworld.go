package main

import (
	"fmt"
	"reflect"
	"time"
)

// variable
var a string = "asad"
var varNumber int32 = 10
var varBool bool = true
var varFloat float32 = 1.10
var abc string

// var b, c, d, e = 1, 2, 3, 4 // multiple variable

// var (
// 	a int
// 	b string
// 	c bool = false
// ) //  Variable Declaration in a Block

const abd = "abd constant"

func main() {
	abc = "adasd"
	test1 := 15.00048548 // infer type
	fmt.Println("Hello World!")
	fmt.Println("Hello World!", time.Now())
	fmt.Println(a)
	fmt.Println(varNumber)
	fmt.Println(varBool)
	fmt.Println(varFloat)
	fmt.Println(reflect.TypeOf(test1))
	fmt.Println(abc)
	// fmt.Println(b)
	fmt.Println(abd)

	// array

	var arr1 = [3]int{1, 2, 3}
	arr2 := [2]string{"asad", "santo"}
	fmt.Println(arr1)
	arr2[0] = "rename asad"
	fmt.Println(len(arr2))
	var arr3 = [...]int{1, 2, 3, 4}
	fmt.Println(arr3)
}
