// [[file:../README.org::sort-colors][sort-colors]]
package main

import (
	"fmt"
	"log"
)

func swap(nums []int, i, j int) {
	tempNum := nums[i]
	nums[i] = nums[j]
	nums[j] = tempNum
}

// sortColors() is bubble sort implementation for colors.
//
// Process: Swap, check order, one-round bubbling, repeated n times.
// Runtime: 4 ms, faster than 25.87% of Go online submissions for Sort Colors.
// Memory Usage: 2.1 MB, less than 67.82% of Go online submissions for Sort Colors.
//
// ♥ 19:56:37.689437 | nums: [2 0 2 1 1 0]
// ♥ 19:56:37.689492 | Output: [0 0 1 1 2 2]
func sortColors(nums []int) []int {
	len := len(nums)
	// loop j repeats i-loop n times
	for j := 0; j < len; j++ {
		// loop i for each len(nums)
		for i := 0; i < len-1; i++ {
			if nums[i] > nums[i+1] {
				swap(nums, i, i+1)
			}
		}
	}
	return nums
}

func main() {
	log.SetPrefix("\n♥ ") // SetPrefix sets the output prefix for the standard logger\.
	log.SetFlags(4)       // SetFlags sets the output flags for the standard logger\. The flag bits are Ldate, Ltime, and so on\.
	fmt.Println("0075-sort-colors/main.go: main()")
	nums := []int{2, 0, 2, 1, 1, 0}
	log.Println("| nums:", nums)
	output := sortColors(nums)
	log.Println("| Output:", output)
}
// sort-colors ends here
