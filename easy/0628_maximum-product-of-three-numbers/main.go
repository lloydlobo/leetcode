// [[file:../README.org::maximum-product-of-three-numbers][maximum-product-of-three-numbers]]
//
// 628. Max imum Product of Three Numbers.
//
// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.
//
// Example 1:
// Input: nums = [1,2,3]
// Output: 6
//
// Example 2:
// Input: nums = [1,2,3,4]
// Output: 24
//
// Example 3:
// Input: nums = [-1,-2,-3]
// Output: -6
//
// Constraints:
// - 3 <= nums.length <= 104
// - -1000 <= nums[i] <= 1000
package main

import (
	"github.com/lloydlobo/leetcode/singlescan"
	"github.com/lloydlobo/leetcode/sorting"
	"github.com/lloydlobo/leetcode/testcases"
)

func main() {
	arrNums, arrWant := testcases.GetMainArgs()
	n := len(arrNums) - 1

	testcases.ExecForLoop(singlescan.MaximumProductSingleScan, arrNums, arrWant, n)
	testcases.ExecForLoop(sorting.MaximumProductSorting, arrNums, arrWant, n)
}

// maximum-product-of-three-numbers ends here
