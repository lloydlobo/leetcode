// [[file:../README.org::binary-search][binary-search]]
//
// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// https://leetcode.com/problems/binary-search/
//
// Example 1:
//
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:
//
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
//
// Example 3:
//
// Input: arr[] = {10, 20, 30, 50, 60, 80, 110, 130, 140, 170}, x = 110
// Output: 6
// Explanation: Element x is present at index 6.
//
// Example 4:
//
// Input: arr[] = {10, 20, 30, 40, 60, 110, 120, 130, 170}, x = 175
// Output: -1
// Explanation: Element x is not present in arr[].
//
// Constraints:
//
//   - 1 <= nums.length <= 104
//   - -104 < nums[i], target < 104
//   - All the integers in nums are unique.
//   - nums is sorted in ascending order.

package main

import (
	"fmt"
	"log"
)

// Runtime: 68 ms, faster than 27.39% of Go online submissions for Binary Search.
// Memory Usage: 6.7 MB, less than 94.45% of Go online submissions for Binary Search.
func Search(nums []int, target int) int {
	var output int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			output = i
			break
		} else {
			output = -1
		}
	}

	return output
}

func main() {
	log.Println("| 0704_binary-search/main.go: main()")
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	output := Search(nums, target)
	fmt.Println("\nOutput:", output)
}

// binary-search ends here
