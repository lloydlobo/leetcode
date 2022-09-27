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

// Log Benchmarks:
//
// 2022-09-27 12:39
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifference
// BenchmarkMaxProductDifference-12          331869              3434 ns/op
// PASS
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 1.18
//
// 2022-09-27 12:38
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifference-12         337892       3349 ns/op
// PASS
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 1.174s
