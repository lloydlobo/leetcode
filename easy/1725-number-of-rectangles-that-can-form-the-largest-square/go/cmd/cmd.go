package cmd

import (
	"math"
	"sort"

	"github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/go/testcases"
)

// countGoodRectangles returns an int.
// rect[4,6] => sqr[4,4]
//
// Input: rectangles = [[5,8],[3,9],[5,12],[16,5]]
// Output: 3
// Explanation: The largest squares you can get
// from each rectangle are of lengths [5,3,5,5].
// The largest possible square is of length 5,
// and you can get it out of 3 rectangles.
// Example 2:
// Input: rectangles = [[2,3],[3,7],[4,3],[3,7]]
// Output: 3
//
// Constraints:
// 1 <= rectangles.length <= 1000
// rectangles[i].length == 2
// 1 <= li, wi <= 109
// li != wi
//
// rectangles[i] = [li, wi]
// length li | width wi
//
// if k <= li && k<= wi
//
//	Cut the ith rectangle to form a square of side len = k.
//
// if li % wi = 0 || wi % li = 0
//
//	maxLen := side length OR value of largest square possible.
//
// Runtime: 30 ms |	Memory: 6.5 MB
func CountGoodRectangles(rectangles [][]int) int {
	var (
		maxLenS []int
		counter int
	)
	count := len(rectangles)
	// For each rectangles.
	for i := 0; i < count; i++ {
		li, wi := float64(rectangles[i][0]), float64(rectangles[i][1]) // Get length & width of each rect.
		maxLen := math.Min(li, wi)                                     // Let maxLen for each sqare be the, minimum val between length & width.
		maxLenS = append(maxLenS, int(maxLen))                         // Collect maxLens for each rect.
	}

	sort.Ints(maxLenS)            // Sort the collected maxLens in ascending order.
	chosenOne := maxLenS[count-1] // The one we want is the biggest of them all.
	// count occurence of 3, count occurence of 5.
	for i := 0; i < count; i++ {
		if maxLenS[i] == chosenOne {
			counter++
		}
	}
	// Return the number of rectangles that can make a
	// square with a side length of maxLen.
	return counter
}

func CountGoodRectanglesCompare(rectangles [][]int) int {
	maxLen := 0
	counter := 0
	k := 0

	for _, rectangle := range rectangles {
		if rectangle[0] > rectangle[1] {
			k = rectangle[1]
		} else {
			k = rectangle[0]
		}
		if maxLen < k {
			maxLen = k
			counter = 1
		} else if maxLen == k {
			counter++
		}
	}
	return counter
}

func CountGoodRectanglesMap(rectangles [][]int) int {
	array := make(map[int]int, 0)
	k := 0

	for _, rectangle := range rectangles {
		switch {
		case rectangle[0] > rectangle[1]:
			k = rectangle[1]
		default:
			k = rectangle[0]
		}
		array[k]++
	}
	max := 0
	for key, _ := range array {
		if max < key {
			max = key
		}
	}
	return array[max]
}

/*
1725. Number Of Rectangles That Can Form The Largest Square

	You are given an array rectangles where rectangles[i] = [li, wi]
	represents the ith rectangle of length li and width wi.
	You can cut the ith rectangle to form a square with a side length of k
	if both k <= li and k <= wi. For example, if you have a rectangle [4,6],
	you can cut it to get a square with a side length of at most 4.
	Let maxLen be the side length of the largest square you
	can obtain from any of the given rectangles.

	Return the number of rectangles that can make a
	square with a side length of maxLen.

	Example 1:
	Input: rectangles = [[5,8],[3,9],[5,12],[16,5]]
	Output: 3
	Explanation: The largest squares you can get
	from each rectangle are of lengths [5,3,5,5].
	The largest possible square is of length 5,
	and you can get it out of 3 rectangles.
	Example 2:
	Input: rectangles = [[2,3],[3,7],[4,3],[3,7]]
	Output: 3

	Constraints:
	1 <= rectangles.length <= 1000
	rectangles[i].length == 2
	1 <= li, wi <= 109
	li != wi
*/

// r = append(r, liwi0, liwi1, liwi2, liwi3)
func Execute() []int {
	var output []int

	arr, n := testcases.GetTestcases()

	for i := 0; i < n; i++ {
		output = append(output, CountGoodRectangles(arr[i]))
	}

	return output
}
