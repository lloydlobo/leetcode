package main

import (
	"fmt"
	"log"
	"time"
	// "time"
)

func swp(arr *[]int, idx1, idx2 int) {
	temp := (*arr)[idx2]
	(*arr)[idx2] = (*arr)[idx1]
	(*arr)[idx1] = temp
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
// Memory Usage: 2.1 MB, less than 67.82% of Go online submissions for Sort Colors.
func sortColors(nums []int) {
	// quicksort(&nums. 0, len(nums)-1)

	// One pass with constant extra space
	// using 3 pointer.
	l, r := 0, len(nums)-1
	i := 0
	for i <= r {
		if nums[i] == 0 {
			swp(&nums, l, i)
			l++
		} else if nums[i] == 2 {
			swp(&nums, r, i)
			r--
			i--
		}
		i++
	}
}

func main() {
	log.SetFlags(4)
	log.SetPrefix("")

	nums := []int{2, 0, 0, 1, 2, 0, 0, 1, 0, 2, 1, 0, 2, 1, 2, 2, 1, 0, 1, 2, 1, 2, 1, 2, 0, 0, 1, 2, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	log.Printf("%v\n\n", nums)
	// log.SetOutput()

	timeStart := time.Now()
	sortColors(nums)
	timeDiff := time.Now().Sub(timeStart)

	log.Printf("%v\n\n", nums)
	fmt.Printf("\nProcess exited in %s.", timeDiff)
	// log.Output(int(timeDiff), "Comepleted!")
}

// func partition(array *[]int, p, r int) int {
//     q:=p
//     for i:=p; i < r; i++ {
//         if (*array)[i] < (*array)[r] {
//             swap(array, i, q)
//             q++
//         }
//     }
//     swap(array, q, r)
//     return q
// }

// func quicksort(array *[]int, p, r int) {
//     if p < r {
//         q := partition(array, p, r)
//         quicksort(array, p, q-1)
//         quicksort(array, q+1, r)
//     }
// }
