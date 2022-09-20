// [[file:../README.org::check-if-array-is-sorted-and-rotated][check-if-array-is-sorted-and-rotated]]
package main

import (
	"time"

	testcase "github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated/testcase"
)

func sleep(x time.Duration) {
	time.Sleep(time.Millisecond * x)
}

func main() {
	go sleep(0)
	arrNums, arrWants := testcase.GetTestcase()
	testcase.ExecForLoop(Check, arrNums, arrWants, len(arrNums))
	testcase.ExecForLoop(CheckMiss, arrNums, arrWants, len(arrNums))
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Check if Array Is Sorted and Rotated.
// Memory Usage: 2 MB, less than 87.88% of Go online submissions for Check if Array Is Sorted and Rotated.
func CheckMiss(nums []int) bool {
	miss := false
	for i := 1; i < len(nums); i++ {
		prevMoreThanCurr := nums[i-1] > nums[i]
		if !miss && prevMoreThanCurr {
			miss = true
		} else if prevMoreThanCurr {
			return false
		}
	}
	if miss {
		isLastLessThanEqualFirst := nums[len(nums)-1] <= nums[0]
		return isLastLessThanEqualFirst
	}
	return true
}

// if array is sorted and rotated then, there is only 1 break point where (nums[x] > nums[x+1]),
// if array is only sorted then, there is 0 break point.
//
// Compare all neignbour elements (a,b) in nums,
// the case of a > b can happen at most once.
// Note that the first element and the last element are also connected.
// If all a <= b, nums is already sorted.
// If all a <= b but only one a > b,
// rotate and make b the first element.
// Otherwise return false.
// # Complexity
// Time O(n)
// Space O(1)
// Runtime: 0 ms.
// Memory Usage: 2.1 MB.
func Check(nums []int) bool {
	count, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true // return (count <= 1)
}

// swap
//
// swapF := reflect.Swapper(nums)
// swapF(nums[i%count], nums[(i+1)%count])
// Swapper returns a function that swaps the elements in the provided slice.
// Swapper panics if the provided interface is not a slice.
//
// 1752. Check if Array Is Sorted and Rotated
// Easy
// Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.
// There may be duplicates in the original array.
//
// Note: An array A rotated by x positions results in an array B of the same length such that A[i] == B[(i+x) % A.length], where % is the modulo operation.
//
// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
/*
# Examples:

    Example 1:

    Input: nums = [3,4,5,1,2]
    Output: true
    Explanation: [1,2,3,4,5] is the original sorted array.
    You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].

    Example 2:

    Input: nums = [2,1,3,4]
    Output: false
    Explanation: There is no sorted array once rotated that can make nums.

    Example 3:

    Input: nums = [1,2,3]
    Output: true
    Explanation: [1,2,3] is the original sorted array.
    You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.
*/

/*
	featureVector := []interface{}{
		[]int{1, 2}, []float64{1.2, 2.2}, []string{"a", "b"},
	}
	featureVector = append(featureVector, []byte{'x', 'y'})
	fmt.Printf("%#v", featureVector)

	// []interface{}{[]int{1, 2}, []float64{1.2, 2.2}, []string{"a", "b"}, []uint8{0x78, 0x79}}
*/

// check-if-array-is-sorted-and-rotated ends here
