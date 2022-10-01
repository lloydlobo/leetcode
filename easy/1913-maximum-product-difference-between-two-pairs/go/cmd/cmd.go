package cmd

import (
	"log"

	algoClone "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/algo/clone"
	algoMinmax "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/algo/minmax"
	algoSort "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/algo/sort"
	testcases "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/testcases"
)

func ExecuteFunction(f func(n []int) int, nums []int, want int) (int, int) {
	var got int
	for i := 0; i < len(nums); i++ {
		got = f(nums)
	}

	log.Printf("got: %v; want: %v\n", got, want)
	return got, want
}

func Execute() {
	arrNums, arrWants := testcases.GetTestcases()

	for i := 0; i < len(arrNums); i++ {
		nums, want := arrNums[i], arrWants[i]

		ExecuteFunction(algoSort.MaxProductDifferenceSort, nums, want)
		ExecuteFunction(algoClone.MaxProductDifferenceClone, nums, want)
		ExecuteFunction(algoMinmax.MaxProductDifferenceMinMaxWithStruct, nums, want)
		ExecuteFunction(algoMinmax.MaxProductDifferenceMinMaxWithoutStruct, nums, want)
		// ExecuteFunction(algoMinmax.MinMax, nums, want)
	}
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
