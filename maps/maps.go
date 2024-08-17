package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	m["name"] = "asad"
	m["lname"] = "santo"

	// delete(m, "lname")
	// clear(m) //? clear map
	fmt.Println(m["name"], m["lname"])
	fmt.Println(len(m))

	ab := map[string]int{"price": 1, "total": 10}
	ab2 := map[string]int{"price": 1, "total": 19}
	fmt.Println(ab)

	val, OK := ab["price"] //? 1st property is value like val and OK is true or false
	fmt.Println(val)
	if OK {
		fmt.Println("all okk")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println("is equal::", maps.Equal(ab, ab2))

}
