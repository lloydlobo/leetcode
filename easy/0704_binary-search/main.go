// [[file:../README.org::binary-search][binary-search]]

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// https://leetcode.com/problems/binary-search/
//
/*
   EXAMPLES:
     Example 1:
        Input: nums = [-1,0,3,5,9,12], target = 9
        Output: 4
        Explanation: 9 exists in nums and its index is 4

     Example 2:
        Input: nums = [-1,0,3,5,9,12], target = 2
        Output: -1
        Explanation: 2 does not exist in nums so return -1

     Example 3:
        Input: arr[] = {10, 20, 30, 50, 60, 80, 110, 130, 140, 170}, x = 110
        Output: 6
        Explanation: Element x is present at index 6.

     Example 4:
        Input: arr[] = {10, 20, 30, 40, 60, 110, 120, 130, 170}, x = 175
        Output: -1
        Explanation: Element x is not present in arr[].
*/
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
	// "time"
)

// Runtime: 68 ms, faster than 27.39% of Go online submissions for Binary SearchForLoop.
// Memory Usage: 6.7 MB, less than 94.45% of Go online submissions for Binary SearchForLoop.
func SearchForLoop(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

func getRandomNumsSlice(count int) ([]int, int, int) {
	nums := []int{10, 20, 30, 50, 60, 80, 110, 130, 140, 170}
	len := len(nums)
	for idx := 0; idx < count; idx++ {
		// nums = append(nums, rand.Int())
		nums = append(nums, nums[len-1]+idx*10)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// len2 := len(nums)
	// mid := len2 - len1
	// fmt.Printf("nums: %v\n", nums)

	return nums, 60, 5

}

func getNumsTargets() ([]int, int, int) {
	nums := []int{10, 20, 30, 50, 60, 80, 110, 130, 140, 170}
	target := 110
	want := 6

	return nums, target, want
}

// Binary Search Algorithm: The basic steps to perform Binary Search are:
//
// https://www.geeksforgeeks.org/binary-search/
//
//	Begin with the mid element of the whole array as a search key.
//	If the value of the search key is equal to the item then return an index of the search key.
//	Or if the value of the search key is less than the item in the middle of the interval, narrow the interval to the lower half.
//	Otherwise, narrow it to the upper half.
//	Repeatedly check from the second point until the value is found or the interval is empty.
//
// # Binary Search Algorithm can be implemented in the following two ways
//
// Iterative Method
// Recursive Method
func SearchIterative(nums []int, target int) int {
	n := len(nums) // length of the slice/array
	left := 0      // left: is the first position at 0
	right := n - 1 // right: 0 index means we subtract one from len

	for right-left > 1 {
		mid := (right + left) / 2
		// time.Sleep(time.Second * 1)
		// fmt.Printf("nums: %2v \n ", nums)
		// fmt.Printf("\rleft: %2v: %v | mid: %2v: %v | right: %2v: %v \n", left, nums[left], mid, nums[mid], right, nums[right])

		if target == nums[mid] {
			return mid // If target is in the middle itself
		} else if target > nums[mid] {
			// increment the position to right
			left = mid + 1 // If target is less than mid then it will be in the right subarray/slice
		} else {
			right = mid - 1 // else the target can only be present in the left subarray/slice
		}
	} // for as a while loop

	// If the element is not present in slice/array
	return -1
}

// SearchRecursive
//
// Credits: https://www.geeksforgeeks.org/binary-search/
//
// l, r := 0, len(nums)-1 // left & right
func SearchRecursive(nums []int, target int, l, r int) int {
	if r >= l {
		mid := (r + l) / 2 // Or l + (r - l) / 2.

		if nums[mid] == target {
			return mid // If the target is in the middle itself.
		} else if nums[mid] > target {
			rSub := mid - 1 // move to the left
			return SearchRecursive(nums, target, l, rSub)
		} else {
			lAdd := mid + 1 // move to the right
			return SearchRecursive(nums, target, lAdd, r)
		}
	}

	return -1
}

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("| 0704_binary-search/main.go: main()")
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

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
