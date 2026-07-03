package main

import (
	"fmt"
	"errors"
)


func int_divide(a, b int) (int, error) {
	var zero int
	if b == zero {
		return 0, errors.New("Can't divide by 0")
	}
	return a/b, nil
}

func divide(a, b float64) (float64, error) {
	var zero float64
	if b == zero {
		return 0, errors.New("Can't divide by 0")
	}
	return a/b, nil
}

func main() {
	a := 1.0; b:= 2.0
	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide result: ", res)
	}
	
	a = 1.0; b= 0.0
	res, err = divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide result: ", res)
	}

	// Ignoring error
	res, _ = divide(a, b)
	fmt.Println("Divide result: ", res)

	
	c := 1; d:= 2
	res2, err := int_divide(c, d)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide result: ", res2)
	}
	
	// Ignoring error
	res2, _ = int_divide(c, d)
	fmt.Println("Divide result: ", res2)
	
	c = 1; d= 0
	res2, err = int_divide(c, d)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide result: ", res2)
	}
}
