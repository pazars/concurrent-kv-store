package main

import "fmt"

func PointInc(val *int) {
	(*val)++
}

func ValInc(val int) int {
	newVal := val + 1
	return newVal
}

func main() {
	a := "Hello, World!"
	fmt.Printf("Value: %s; Address: %v\n", a, &a)

	p := &a // *string
	fmt.Printf("Value: %s; Address: %v\n", *p, p)
	fmt.Printf("Same addresses: %t\n\n", &a == p)

	*p = "Hello, Stanley!"
	fmt.Printf("Value: %s; Address: %v\n", *p, p)
	fmt.Printf("Value: %s; Address: %v\n", a, &a)
	fmt.Printf("Same addresses: %t\n\n", &a == p)

	num := 10

	fmt.Printf("Value: %d\n", num)

	PointInc(&num)
	fmt.Printf("Value: %d\n", num)

	num2 := ValInc(num)
	fmt.Printf("Value1: %d; Value2: %d\n", num, num2)
	
	PointInc(&num2)
	fmt.Printf("Value1: %d; Value2: %d\n", num, num2)
	
}
