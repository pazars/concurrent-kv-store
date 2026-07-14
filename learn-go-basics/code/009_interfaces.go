package main

import "fmt"

// The interface
type Stringer interface {
	String() string
}

// The struct
type StringCol struct {
	s1 string
	s2 string
}

// The method
func (sc *StringCol) String() string {
	s1 := (*sc).s1
	s2 := (*sc).s2
	return fmt.Sprintf("s1: %s; s2: %s", s1, s2)
}

// Function accepting interface as param
func printString(si Stringer) {
	fmt.Println(si.String())
}

func main() {
	sc := StringCol{s1: "abc", s2: "123"}
	st := sc.String()

	// Output is the same
	fmt.Println(st)
	fmt.Println(&sc)
	printString(&sc)
}
