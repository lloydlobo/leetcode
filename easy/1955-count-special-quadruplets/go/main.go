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
	out = algo.ExecuteAlgo1(Nums)
	fmt.Printf("outAlgo1: %v\n", out)
	// out = algo.Execute(Nums)
	out = countQuadruplets(Nums)
	fmt.Printf("outCountQua...: %v\n", out)
}

func countQuadruplets(nums []int) int {
	MAP := make(map[int]int)
	count, n := 0, len(nums)
	for b := 1; b < n; b++ {
		// a + b
		for a := 0; a < b; a++ {
			MAP[nums[a]+nums[b]]++
		}
		//d - c
		c := (b + 1)
		for d := b + 2; d < n; d++ {
			count += MAP[nums[d]-nums[c]]
		}
	}
	return count
}
