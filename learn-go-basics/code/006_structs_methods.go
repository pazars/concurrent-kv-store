package main

import "fmt"

type Counter struct {
	value int
}

func (s *Counter) Inc() {
	s.value++
}

func (s Counter) Value() int {
	return s.value
}

func NewCounter() *Counter {
	return &Counter{}
}

func main() {
	counter := NewCounter()
	fmt.Printf("#1 Counter value: %d\n", counter.Value())

	counter.Inc()
	fmt.Printf("#2 Counter value: %d\n", counter.Value())

	fmt.Printf("Counter: %+v\n", counter)
}
