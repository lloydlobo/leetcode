package main

import (
	"testing"

	algoClone "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/algo/clone"
	algoMinmax "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/algo/minmax"
	algoSort "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/algo/sort"
	testcases "github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go/testcases"
)

func TestRunMain(t *testing.T) {
	main()
}

func TestMaxProductDifferenceMinmax(t *testing.T) {
	nums, wants := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num, want := nums[i], wants[i]

		got := algoMinmax.MaxProductDifferenceMinMax(num)
		if got != want {
			t.Errorf("MaxProductDifference(%v) = %v; want: %v", num, got, want)
		}
	}
}

func BenchmarkMaxProductDifferenceMinmax(b *testing.B) {
	nums, _ := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := 0; j < b.N; j++ {
			algoMinmax.MaxProductDifferenceMinMax(num)
		}
	}
}
func TestMaxProductDifferenceSort(t *testing.T) {
	nums, wants := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num, want := nums[i], wants[i]

		got := algoSort.MaxProductDifferenceSort(num)
		if got != want {
			t.Errorf("MaxProductDifference(%v) = %v; want: %v", num, got, want)
		}
	}
}

func BenchmarkMaxProductDifferenceSort(b *testing.B) {
	nums, _ := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := 0; j < b.N; j++ {
			algoSort.MaxProductDifferenceSort(num)
		}
	}
}

func TestMaxProductDifferenceClone(t *testing.T) {
	nums, wants := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num, want := nums[i], wants[i]

		got := algoClone.MaxProductDifferenceClone(num)
		if got != want {
			t.Errorf("MaxProductDifference(%v) = %v; want: %v", num, got, want)
		}
	}
}

func BenchmarkMaxProductDifferenceClone(b *testing.B) {
	nums, _ := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := 0; j < b.N; j++ {
			algoClone.MaxProductDifferenceClone(num)
		}
	}
}

// Log Benchmarks:
//
// 2022-09-27 13:19
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifferenceMinmax
// BenchmarkMaxProductDifferenceMinmax-12          52780066                22.92 ns/op
// BenchmarkMaxProductDifferenceSort-12             5168884               234.9 ns/op
// BenchmarkMaxProductDifferenceClone-12             348068              3392 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 3.907s
//
// 2022-09-27 12:55
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifference-12                 5090146               227.5 ns/op
// BenchmarkMaxProductDifferenceClone-12             338812              3354 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 2.577s
//
// 2022-09-27 12:51
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifference-12                 5174109               229.7 ns/op
// BenchmarkMaxProductDifferenceClone-12             337312              3439 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 2.62
//
// 2022-09-27 12:39
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifference-12          331869              3434 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 1.18
//
// 2022-09-27 12:38
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifference-12         337892       3349 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 1.174s
