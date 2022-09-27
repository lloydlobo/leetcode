package main

import (
	"testing"

	cmd "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/cmd"
	testcases "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/testcases"
)

func TestRunMain(t *testing.T) {
	main()
}

func TestMaxProductDifference(t *testing.T) {
	nums, wants := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num, want := nums[i], wants[i]

		got := cmd.MaxProductDifferrence(num)
		if got != want {
			t.Errorf("MaxProductDifference(%v) = %v; want: %v", num, got, want)
		}
	}
}

func BenchmarkMaxProductDifference(b *testing.B) {
	nums, _ := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := 0; j < b.N; j++ {
			cmd.MaxProductDifferrence(num)
		}
	}
}
