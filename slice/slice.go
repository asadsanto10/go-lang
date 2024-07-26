package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice
	// mySlice := [5]int{1, 2, 3, 4}
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	// var nums []int

	// nums = make([]int, 2)
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// fmt.Println(cap(nums))

	// var nums = make([]int, 0, 3)
	// capacity -> maximum number of elements
	// fmt.Println(cap(nums))

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 7)
	// nums = append(nums, 7)

	// fmt.Println("nums arrays::", nums)
	// fmt.Println("capicity::", cap(nums))
	// fmt.Println("len::", len(nums))

	//? copy functions

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	//? slice operatior
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(nums[0:3])
	fmt.Println(nums[0:3])
	fmt.Println(nums[1:])

	// slice
	var numst = []int{1, 2, 3, 4, 5}
	var numst2 = []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Equal(numst, numst2))

	var twoDnums = [][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	fmt.Println(twoDnums)
}
