package algo

import (
	"fmt"
	"sort"
	"time"
)

// var Nums = []int{2, 16, 9, 27, 3, 39} // 2; Got 1
// [9,6,8,23,39,23] // 2; Got 1
func LeetCountQuadruplets(nums []int) int {
	// mp := SortLess(nums)
	sort.Ints(nums)
	m := GetMap(nums)
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("m: %v\n", m)
	counter, n := 0, len(m)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				// var Nums = []int{1, 1, 1, 3, 5} // 4
				if i <= j && j < k || i < j && j <= k || i == j && j == k {
					if m[i]+m[j]+m[k] == m[k+1] {
						counter++
					}
				}
			}
		}
	}
	return counter
}

// index: 0 + 1 + 2 + 3
// rhs := fourth element in each loop
// lhs := i + j + k = rhs
// where i < j < k
// MAP: map[0:1 1:1 2:1 3:3 4:5]
func ProcessMap(m map[int]int) int {
	counter, n := 0, len(m)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j; k < n; k++ {
				if i < j && j < k {
					if m[i]+m[j]+m[k] == m[k+1] {
						counter++
					}
				}
			}
		}
	}
	return counter
}

// Input: nums = [1,1,1,3,5]; Output: 4
// Explanation: The 4 quadruplets that satisfy the requirement are:
// - (0, 1, 2, 3): 1 + 1 + 1 == 3
// - (0, 1, 3, 4): 1 + 1 + 3 == 5
// - (0, 2, 3, 4): 1 + 1 + 3 == 5
// - (1, 2, 3, 4): 1 + 1 + 3 == 5
func GetMap(nums []int) (m map[int]int) {
	m = make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		m[i] = nums[i]
	}
	return m
}
func Execute(nums []int) int { return CountQuadruplets(nums) }

// - Given a 0-indexed integer array nums,
// return the number of distinct quadruplets (a, b, c, d) such that:
func CountQuadruplets(nums []int) (out int) {
	n := SortLess(nums)
	MAP := GetMap(n)
	PROCESS_MAP := ProcessMap(MAP)
	fmt.Printf("MAP: %v\n", MAP)
	fmt.Printf("PROCESS_MAP: %v\n", PROCESS_MAP)

	out = LastIsEqualToSumOfAll(n)
	return out
}

// nums[a] + nums[b] + nums[c] == nums[d], and
// a < b < c < d
func SortLess(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	fmt.Println(l, r)

	for l <= r {
		mid := l + (r-l)/2
		fmt.Printf("mid: %v\n", mid)
		fmt.Println(l, r)

		// time.Sleep(time.Millisecond * 1000)
		if nums[mid] < nums[l] {
			SwapAWithB(&nums, l, mid) // nums[left], nums[mid] = nums[mid], nums[left]
			l++
		} else if nums[mid] > nums[r] {
			SwapAWithB(&nums, r, mid) // nums[left], nums[mid] = nums[mid], nums[left]
			r--
		} else {
			l++
		}
	}

	return nums
}

func SwapAWithB(n *[]int, a, b int) {
	(*n)[a], (*n)[b] = (*n)[b], (*n)[a]
}

func GetLen(n []int) int {
	return len(n)
}

/*
	** Explanation

https://leetcode.com/problems/count-special-quadruplets/

  - Given a 0-indexed integer array nums,
    return the number of distinct quadruplets (a, b, c, d) such that:

    #+begin_example
    nums[a] + nums[b] + nums[c] == nums[d], and
    a < b < c < d
    #+end_example

*** Constraints:

	4 <= nums.length <= 50
	1 <= nums[i] <= 100

*** Examples
#+begin_example
Example 1:

Input: nums = [1,2,3,6]
Output: 1
Explanation: The only quadruplet that satisfies the requirement is (0, 1, 2, 3) because 1 + 2 + 3 == 6.

Example 2:

Input: nums = [3,3,6,4,5]
Output: 0
Explanation: There are no such quadruplets in [3,3,6,4,5].

Example 3:

Input: nums = [1,1,1,3,5]
Output: 4
Explanation: The 4 quadruplets that satisfy the requirement are:
- (0, 1, 2, 3): 1 + 1 + 1 == 3
- (0, 1, 3, 4): 1 + 1 + 3 == 5
- (0, 2, 3, 4): 1 + 1 + 3 == 5
- (1, 2, 3, 4): 1 + 1 + 3 == 5

#+end_example
*/
func LastIsEqualToSumOfAll(nums []int) int {
	var (
		sumMap  = make(map[int]int)
		n       = len(nums)
		counter = 0
	)
	time.Sleep(time.Millisecond * 1000)

	for i := 1; i < 4; i++ {
		for j := i; j < n-1; j++ {
			time.Sleep(time.Millisecond * 1000)
			a, b := nums[j-1], nums[(j)%n]
			c, d := nums[(j+1)%n], nums[(j+2)%n]
			if a+b+c == d {
				sumMap[i+j] = nums[j]
				counter++
			}
		}
	}
	fmt.Printf("sumMap: %v\n", sumMap)
	return len(sumMap)
}

func ArchiveAlgo0Loop(n, counter, left, right, next int, nums, sum []int) {
	// var (
	// 	n     = len(nums)
	// 	last  = nums[n-1]
	// 	right = nums[0]
	// 	left  = 0
	// 	next  = 0
	// 	sum   = []int{}
	// )
	// counter := 4
	// return last == next-left
	for j := n - 1; j > 0; j-- {
		if counter >= 0 {
			counter--
			fmt.Printf("counter: %v\n", counter)
			fmt.Printf("num: %v\n", nums[j])
			right = nums[j]
			left = nums[j-1]
			fmt.Println(left, right)
			next = right + left
			sum = append(sum, next)
			fmt.Printf("next: %v\n", next)
		} else {
			fmt.Printf("sumNext: %v\n", next)
		}
	}

}
