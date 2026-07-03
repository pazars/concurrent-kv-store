package main

import "fmt"

func main() {
	var num int
	var str string
	var bo bool
	var arr []int

	fmt.Printf("Empty integer: %d\n", num)
	fmt.Printf("Empty string: w/ %%s %s and w/ %%v %v\n", str, str)
	fmt.Printf("Empty boolean: %t\n", bo)
	fmt.Printf("Empty integer array: w/ %%d %d, w/ %%v %v, and %%+v %+v\n", arr, arr, arr)

	const MAX_VALUE int = 5
	const MAXX_VALUE int = 3

	fmt.Printf("Max values: %d, %d\n", MAX_VALUE, MAXX_VALUE)

	const (
		Red = iota
		Blue
		Green
	)

	fmt.Println(Red, Green)
}
