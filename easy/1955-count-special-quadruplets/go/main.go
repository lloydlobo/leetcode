package main

import (
	"fmt"

	"github.com/lloydlobo/leetcode/easy/1955-count-special-quadruplets/go/algo"
)

// var Nums = []int{1, 2, 3, 6} // 1
// var Nums = []int{3, 3, 6, 4, 5} // 0
var Nums = []int{1, 1, 1, 3, 5} // 4
// var Nums = []int{2, 16, 9, 27, 3, 39} // 2; Got 1
// var Nums = []int{9, 6, 8, 23, 39, 23} // 2; Got 1
func main() {
	out := algo.LeetCountQuadruplets(Nums)
	fmt.Printf("outLeet: %v\n", out)
	// out = algo.Execute(Nums)
	// fmt.Printf("out: %v\n", out)
}
