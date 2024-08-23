package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Hello, World!\n")

	lst := []int{1, 3, 2, 4}
	fmt.Printf("lst: %v\n", lst)
	sort.Slice(lst, func(i, j int) bool {
		return i > j
	})
	fmt.Printf("lst: %v\n", lst)

	x := 4

	for x > 2 {
		x--
		fmt.Printf("x: %v\n", x)
	}
}
