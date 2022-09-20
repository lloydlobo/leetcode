package singlescan

import "math"

// MaximumProductSingleScan()
//
// Runtime: 37 ms.
// Memory Usage: 6.5 MB.
// https://leetcode.com/submissions/detail/803690509/
//
// BenchmarkMaximumProductSingleScan-12            145782132                8.209 ns/op
//
// Complexity Analysis:
// Time complexity : O(n).Only one iteration over the nums array of length nnn is required.
// Space complexity : O(1). Constant extra space is used.
//
// We need not necessarily sort the given nums array to find the maximum product.
//
// Instead, we can only find the required 2 smallest values(min1 and min2)
// and the three largest values(max1,max2,max3) in the nums array,
// by iterating over the nums array only once.
//
// At the end, again we can find out the larger value out of
// min1×min2×max1 times max1×max2×max3 to find the required maximum product.
func MaximumProductSingleScan(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n <= min1 {
			min2 = min1
			min1 = n
		} else if n <= min2 {
			min2 = n
		}
		if n >= max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n >= max2 {
			max3 = max2
			max2 = n
		} else if n >= max3 {
			max3 = n
		}
	}
	return int(math.Max(float64(min1*min2*max1), float64(max1*max2*max3)))
}
