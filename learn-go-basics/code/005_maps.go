package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["A"] = 1

	// Key exists
	key := "A"
	v, ok := m[key]
	if !ok {
		fmt.Printf("Couldn't find key in map\n")
	} else {
		fmt.Printf("Found value: %v\n", v)
	}

	// Key doesn't exist
	key = "B"
	v2, ok2 := m[key]
	if !ok2 {
		fmt.Printf("Couldn't find key in map\n")
	} else {
		fmt.Printf("Found value: %v\n", v2)
	}

	m[key] = 2
	for kr, vr := range m {
		fmt.Printf("Key: %v; Value: %v\n", kr, vr)
	}

	// Delete a key
	delete(m, key)

	// nil-map panic
	var n map[string]int

	// read is ok
	v3, ok3 := n[key]

	if !ok3 {
		fmt.Printf("Couldn't find key in map\n")
	} else {
		fmt.Printf("Found value: %v\n", v3)
	}

	// Write is nil-map panic
	// n["Foo"] = 10
}
