package cmd

import (
	"fmt"
	"log"

	testcases "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/testcases"
)

func ExecuteFunction(f func(n []int) int, nums []int, want int) (int, int) {
	var got int
	for i := 0; i < len(nums); i++ {
		got = f(nums)
	}
	fmt.Println()
	log.Printf("got: %v; want: %v\n", got, want)
	return got, want
}

func Execute() {

	arrNums, arrWants := testcases.GetTestcases()

	for i := 0; i < len(arrNums); i++ {
		nums, want := arrNums[i], arrWants[i]
		// ExecuteFunction(MaxProductDifference, nums, want)
		ExecuteFunction(MaxProductDifferrence, nums, want)
	}
}
