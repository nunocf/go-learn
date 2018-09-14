package main

import "fmt"

func main() {
	s := newSlice(10)

	for _, value := range s {
		if value%2 == 0 {
			fmt.Printf("%v is even\n", value)
		} else {
			fmt.Printf("%v is odd\n", value)
		}
	}
}

func newSlice(size int) []int {
	s := make([]int, size+1)

	for i := range s {
		s[i] = i
	}

	return s
}
