// [[file:../README.org::binary-search][binary-search]]
// Given an
 array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// https://leetcode.com/problems/binary-search/
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
	"sort"
	"stdout"
)

// Runtime: 31 ms, faster than 95.31% of Go online submissions for Binary Search.
// Memory Usage: 6.7 MB, less than 94.48% of Go online submissions for Binary Search.
//
// Runtime: 32 ms, faster than 93.68% of Go online submissions for Binary Search.
// Memory Usage: 6.5 MB, less than 99.96% of Go online submissions for Binary Search.
func searchDivideAndConquer(nums []int, target int) int {
	low, high := 0, len(nums)-1 // left: is the first position at 0// right: 0 index means we subtract one from len
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid // If target is in the middle itself
		} else if nums[mid] > target {
			high-- // high = mid - 1 //  else the target can only be present in the left subarray/slice
		} else {
			low++ // low = low + 1 // If target is less than mid then it will be in the right subarray/slice
		}
	}
	return -1
}

// Runtime: 40 ms, faster than 72.72% of Go online submissions for Binary Search.
// Memory Usage: 7.1 MB, less than 53.18% of Go online submissions for Binary Search.
//
// Runtime: 33 ms, faster than 91.94% of Go online submissions for Binary Search.
// Memory Usage: 7 MB, less than 61.39% of Go online submissions for Binary Search.
func searchBruteForce(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("| 0704_binary-search/main.go: main()")
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	searchRunMain()

	output := SearchForLoop(nums, target)
	fmt.Printf("\nSearchForLoop:%2v\n", output)

	nums, target, _ = getNumsTargets()
	outputIterative := SearchIterative(nums, target)
	fmt.Println("\nSearchIterative:", outputIterative)

	// nums, target, _ = getRandomNumsSlice(20)
	// outputIterative = SearchIterative(nums, target)
	// fmt.Println("\nSearchIterative:", outputIterative)

	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 3
	// nums, target, _ = getRandomNumsSlice(20)
	outputRecursive := SearchRecursive(nums, target, 0, len(nums))
	fmt.Println("\nSearchRecursive:", outputRecursive)

	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 1
	// nums, target, _ = getRandomNumsSlice(20)
	outputRecursive = SearchRecursive(nums, target, 0, len(nums))
	fmt.Printf("\nSearchRecursive: %2v \n\n", outputRecursive)
	// stdout.OutputToFile("file.log")

	stdout.OutputToFile("")

}
// binary-search ends here
