// 944. Delete Columns to Make Sorted
//
// You are given an array of n strings strs, all of the same length.
// The strings can be arranged such that there is one on each line, making a grid. For example, strs = ["abc", "bce", "cae"] can be arranged as:
// abc
// bce
// cae
//
// You want to delete the columns that are not sorted lexicographically. In the above example (0-indexed), columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted while column 1 ('b', 'c', 'a') is not, so you would delete column 1.
// Return the number of columns that you will delete.
//
// Constraints:
// - n == strs.length
// - 1 <= n <= 100
// - 1 <= strs[i].length <= 1000
// - strs[i] consists of lowercase English letters.
package min_deletion_size

import (
	"fmt"
	"strings"
	"time"
)

// cba
// daf
// ghi
func Execute() {
	var (
		out   int
		strs0 = []string{"cba", "daf", "ghi"}
		strs1 = []string{"a", "b"}
		strs2 = []string{"zyx", "wvu", "tsr"}
		strs3 = []string{"rrjk", "furt", "guzm"}
	)
	fmt.Printf("\n%v\n0944: Execute()\n", time.Now().Local())

	out = MinDeletionSize(strs0)
	fmt.Printf("strs0: %v; %v\n", out, out == 1)
	out = MinDeletionSize(strs1)
	fmt.Printf("strs1: %v; %v\n", out, out == 0)
	out = MinDeletionSize(strs2)
	fmt.Printf("strs2: %v; %v\n", out, out == 3)
	out = MinDeletionSize(strs3)
	fmt.Printf("strs3: %v; %v\n", out, out == 2)
}

// MinDeletionSize()
// ~ ~ C C C C
// ~ ~ 0 1 2 3
// R 0 c b a .
// R 1 d a f .
// R 2 g h i .
// R 3 . . . .
func MinDeletionSize(strs []string) int {
	var (
		collect        []string
		countDeleteIdx int   // Counter that tracks unsorted columns to remove.
		idxDel         int   // The index of col to be deleted.
		arrIdxDel      []int // The index of col to be deleted.
		sizeStrs       = len(strs)
		isSorted       bool
	)

	// count occurence of two unsorted instance
	for i := 0; i < sizeStrs; i++ {
		split := strings.Split(strs[i], "") // [c b a] [d a f] [g h i]
		collect = append(collect, split...) // [c b a d a f g h i]
	}

	sizeCollect := len(collect)
	var tempCol = make([]string, 0)
	var arrCols []string

	// Get chars at each count position
	// [c b a d a f g h i] => [c d g b a h a f i]
	for j := 0; j < sizeStrs; j++ {
		for i := 0; i < sizeCollect; i++ {
			if i%sizeStrs == 0 {
				tempCol = append(tempCol, collect[i+j])
			}
		} // end of for loop i.

		isSorted, _ = isUnsortedTwice(tempCol)
		if !isSorted {
			countDeleteIdx++
			idxDel = j
			arrIdxDel = append(arrIdxDel, idxDel)
		}
		fmt.Printf("idxDel: %v\n", idxDel)

		arrCols = append(arrCols, tempCol...) // Append tempCol at each count'th position.
		tempCol = []string{}                  // Clear tempCol.
	} // end of for loop j.

	// Return count of unsorted columns indexed.
	return len(arrIdxDel)
}

// isUnsortedTwice() finds if a slice has 2 points of failure.
// while checking for sorting.
//
// isSorted (R,C) = (0,0) < (0,1) < (0,2)
//
// ~ ~ C C C C
// ~ ~ 0 1 2 3
// R 0 c b a .
// R 1 d a f .
// R 2 g h i .
// R 3 . . . .
func isUnsortedTwice(s []string) (bool, int) {
	var (
		counter int
		ns      = len(s)
		i       int
	)
	for i = 0; i < ns; i++ {

		if s[i] > s[(i+1)%ns] {
			counter++
		}
		if counter > 1 {
			return false, i
		}
	}
	return true, i
}

/*
 * Examples 1
 *  Input: strs = ["cba","daf","ghi"]
 *  Output: 1
 *  Explanation: The grid looks as follows:
 *    cba
 *    daf
 *    ghi
 *  Columns 0 and 2 are sorted, but column 1 is not,
 *  so you only need to delete 1 column.
 *
 * Examples 2
 *  Input: strs = ["a","b"]
 *  Output: 0
 *  Explanation: The grid looks as follows:
 * 	a
 * 	b
 *  Column 0 is the only column and is sorted,
 *  so you will not delete any columns.
 */
