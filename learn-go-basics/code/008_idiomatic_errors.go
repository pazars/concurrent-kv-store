package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findSubstring(sFull, sSub string) (bool, error) {
	lenFull := len(sFull)
	lenSub := len(sSub)
	var slice string

	for i := 0; i <= (lenFull - lenSub); i++ {
		slice = sFull[i : i+lenSub]
		// fmt.Println(slice)
		if slice == sSub {
			return true, nil
		}
	}

	err := fmt.Errorf("substring %s %w", sSub, ErrNotFound)
	return false, err
}

func main() {
	s1 := "Hello, World!"
	s2 := "ld!"
	s3 := "yipee"

	r, err := findSubstring(s1, s2)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("Correctly matched error\n")
	} else {
		fmt.Printf("Returned value: %v\n", r)
	}

	r2, err2 := findSubstring(s1, s3)
	if errors.Is(err2, ErrNotFound) {
		fmt.Printf("Correctly matched error\n")
	} else {
		fmt.Printf("Returned value: %v\n", r2)
	}
}
