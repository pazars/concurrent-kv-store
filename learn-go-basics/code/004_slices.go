package main

import "fmt"

func main() {
	// true
	var arr []int
	fmt.Println(len(arr) == cap(arr))

	// true
	arr2 := make([]int, 3)
	fmt.Println(len(arr2) == cap(arr2))

	// false
	arr3 := make([]int, 0, 1)
	fmt.Println(len(arr3) == cap(arr3))

	arr = append(arr, 1)	
	arr2 = append(arr2, 2)
	arr3 = append(arr3, 3)

	// can be true or false, growth or not decided by runtime allocator
	fmt.Println(len(arr) == cap(arr), len(arr), cap(arr))

	// false (was 3/3, added same backing array (2x growth) with 3 slots but filled just 1)
	fmt.Println(len(arr2) == cap(arr2), len(arr2), cap(arr2))

	// true (was 0/1, adding 1 number results in 1/1)
	fmt.Println(len(arr3) == cap(arr3))

	// [0, 0, 0, 2]
	arr4 := arr2[0:4]
	fmt.Println(arr4)

	// [2, 0]
	arr5 := arr2[3:5]
	fmt.Println(arr5)

	// panic: runtime error: slice bounds out of range
	// arr6 := arr2[3:10]
	// fmt.Println(arr6)

	var arr7 []int
	for _, v := range arr2 {
		arr7 = append(arr7, 2 * (v + 1))
	}

	// [2, 2, 2, 6]
	fmt.Println(arr7)
	// true
	fmt.Println(len(arr7) == cap(arr7))
}
