package main

import (
	"fmt"
	"math/rand"
)

func get_random_name() string {
	names := []string{
		"World",
		"Friend",
		"All",
		"Person",
	}
	num_pick := rand.Intn(len(names))
	return names[num_pick]
}

// This is a single line comment
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// := is declaration + assignment
	// type implicit from function return
	for i:=0; i<3; i++ {
		name := get_random_name()
		greet(name)
	 }
}
