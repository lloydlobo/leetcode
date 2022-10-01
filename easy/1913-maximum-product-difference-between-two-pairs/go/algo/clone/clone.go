// [[file:../../../../README.org::*clone][clone:1]]
package clone

import "sort"

type SortedKeys struct {
	Value int
	Id    int
}

// MaxProductDifferenceClone keeps the nums intact,
// and returns difference of the product of
// the 2 largest and 2 smallest numbers.
//
// Runtime: 69 ms; Memory: 8.5 MB
func MaxProductDifferenceClone(nums []int) int {
	n := len(nums)

	sortedNums := getSortIdx(nums, n)                 // Sort by value while keeping the index beside it.
	lg, sm, _ := getGreatestMultiplier(sortedNums, n) // Multiply and store id of first 2 and last two nums.

	return lg - sm // Return difference of largest to smallest product.
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

func getGreatestMultiplier(n [][]int, size int) (int, int, []int) {
	a, b, y, z := n[0][0], n[1][0], n[size-2][0], n[size-1][0]         // 9 8 4 2
	ida, idb, idy, idz := n[0][1], n[1][1], n[size-2][1], n[size-1][1] // a b y z

	large, small := a*b, y*z             // Populate 0*1, n-1*n. (products)
	abyzIdx := []int{ida, idb, idy, idz} // Populate 0,1,n-1,n. (ids)

	return large, small, abyzIdx
}
// clone:1 ends here
