package cmd

import (
	"sort"
)

// MaxProductDifferrence()
//
// 4 <= nums.length <= 104
// 1 <= nums[i] <= 104
// The product difference between two pairs (a, b) and (c, d)
// is defined as (a * b) - (c * d).
func MaxProductDifferrence(nums []int) int {
	n := len(nums)

	sortedNums := getSortIdx(nums, n)                 // Sort by value while keeping the index beside it.
	lg, sm, _ := getGreatestMultiplier(sortedNums, n) // Multiply and store id of first 2 and last two nums.

	return lg - sm // Return difference of largest to smallest product.
}

func getGreatestMultiplier(nums [][]int, size int) (int, int, []int) {
	a, b, y, z := nums[0][0], nums[1][0], nums[size-2][0], nums[size-1][0]     // 9 8 4 2
	ia, ib, iy, iz := nums[0][1], nums[1][1], nums[size-2][1], nums[size-1][1] // 9 8 4 2

	lg, sm := a*b, y*z               // Populate 0*1, n-1*n. (products)
	lgsmIdx := []int{ia, ib, iy, iz} // Populate 0,1,n-1,n. (ids)

	return lg, sm, lgsmIdx
}

func getSortIdx(nums []int, size int) [][]int {
	var hash = make(map[int][]int)
	var keys = make([][]int, size)
	// Populate hash map with hashed nums and id.
	for i := 0; i < size; i++ {
		// get index of all nums before sorting
		hash[i] = []int{nums[i], i}
	}
	// Populate keys slice with hashed nums and id.
	for i := 0; i < len(hash); i++ {
		keys[i] = []int{hash[i][0], hash[i][1]}
	}
	// Sort by value while keeping the index beside it.
	sort.Slice(keys, func(i, j int) bool {
		return keys[i][0] > keys[j][0]
	})
	return keys
}

/*
m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
	fmt.Println(k, m[k])
}

Output:

Alice 23
Bob 25
Eve 2



*/

/*
Examples.
    Example 1:
        Input: nums = [5,6,2,7,4]
        Output: 34
        Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and indices 2 and 4 for the second pair (2, 4).
        The product difference is (6 * 7) - (2 * 4) = 34.

    Example 2:
        Input: nums = [4,2,5,9,7,4,8]
        Output: 64
        Explanation: We can choose indices 3 and 6 for the first pair (9, 8) and indices 1 and 5 for the second pair (2, 4).
        The product difference is (9 * 8) - (2 * 4) = 64.

*/
/*
 1913. Maximum Product Difference Between Two Pairs
    https://leetcode.com/problems/maximum-product-difference-between-two-pairs/

    The product difference between two pairs (a, b) and (c, d) is defined as (a * b) - (c * d).

    For example, the product difference between (5, 6) and (2, 7) is (5 * 6) - (2 * 7) = 16.

    Given an integer array nums, choose four distinct indices
    w, x, y, and z such that the product difference between pairs
    (nums[w], nums[x]) and (nums[y], nums[z]) is maximized.

    Return the maximum such product difference.

    Constraints:

    4 <= nums.length <= 104
    1 <= nums[i] <= 104
*/
