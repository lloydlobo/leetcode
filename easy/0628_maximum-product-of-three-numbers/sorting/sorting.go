package sorting

import (
	"math"
	"sort"
)

// Sorting Method for maximumProductSorting()
//
// 628. Maximum Product of Three Numbers.
// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.
//
// Find the 3 greatest numbers & Sort them.
// Pick first 3 numbers & Multiply them.
//
// https://leetcode.com/problems/maximum-product-of-three-numbers/solution/
//
// Runtime: 97 ms.
// Memory Usage: 6.8 MB.
//
// BenchmarkMaximumProductSorting-12               29704512                38.49 ns/op
//
// Complexity Analysis:
// Time complexity : O(n log n). Sorting the nums array takes n log n time.
// Space complexity : O(log n). Sorting takes O(log n) space.
//
// ...sort the given numsnumsnums array(in ascending order) and find out the product of the last three numbers.
// But, we can note that this product will be maximum only if all the numbers in numsnumsnums array are positive. But, in the given problem statement, negative elements could exist as well.
// Thus, it could also be possible that two negative numbers lying at the left extreme end could also contribute to lead to a larger product if the third number in the triplet being considered is the largest positive number in the numsnumsnums array.
// Thus, either the product nums[0]×nums[1]×nums[n−1]nums[0] \times nums[1] \times nums[n-1]nums[0]×nums[1]×nums[n−1] or nums[n−3]×nums[n−2]×nums[n−1]nums[n-3] \times nums[n-2] \times nums[n-1]nums[n−3]×nums[n−2]×nums[n−1] will give the required result. Thus, we need to find the larger one from out of these values.
func MaximumProductSorting(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	start := float64(nums[0] * nums[1] * nums[n-1])
	end := float64(nums[n-1] * nums[n-2] * nums[n-3])
	out := int(math.Max(start, end))
	return out
}
