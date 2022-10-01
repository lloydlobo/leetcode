// [[file:../../../../README.org::*sort][sort:1]]
package sort

import (
"sort"
)

// MaxProductDifferenceSort() sorts the slice nums.
// and returns product of first & last two integers.
//
// Using an unstable sort standard function.
//
// https://leetcode.com/submissions/detail/809639232/
//
// Runtime: 34 ms; Memory: 6.4 MB
func MaxProductDifferenceSort(nums []int) int {
n := len(nums)
sort.Ints(nums)
return (nums[n-1] * nums[n-2]) - (nums[0] * nums[1])
}
// sort:1 ends here
