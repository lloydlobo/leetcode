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

func TestMaxProductDifferenceMinmaxWithStruct(t *testing.T) {
	nums, wants := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num, want := nums[i], wants[i]

		got := algoMinmax.MaxProductDifferenceMinMaxWithStruct(num)
		if got != want {
			t.Errorf("MaxProductDifference(%v) = %v; want: %v", num, got, want)
		}
	}
}

func BenchmarkMaxProductDifferenceMinmaxWithStruct(b *testing.B) {
	nums, _ := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := 0; j < b.N; j++ {
			algoMinmax.MaxProductDifferenceMinMaxWithStruct(num)
		}
	}
}

func TestMaxProductDifferenceMinmaxWithoutStruct(t *testing.T) {
	nums, wants := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num, want := nums[i], wants[i]

		got := algoMinmax.MaxProductDifferenceMinMaxWithoutStruct(num)
		if got != want {
			t.Errorf("MaxProductDifference(%v) = %v; want: %v", num, got, want)
		}
	}
}

func BenchmarkMaxProductDifferenceMinmaxWithoutStruct(b *testing.B) {
	nums, _ := testcases.GetTestcases()
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := 0; j < b.N; j++ {
			algoMinmax.MaxProductDifferenceMinMaxWithoutStruct(num)
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
// 2022-09-28 07:20
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifferenceMinmaxWithStruct-12        83828912                14.19 ns/op
// BenchmarkMaxProductDifferenceMinmaxWithoutStruct-12     83060829                14.26 ns/op
// BenchmarkMaxProductDifferenceSort-12                    13768419                87.35 ns/op
// BenchmarkMaxProductDifferenceClone-12                     714735              1538 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 5.783s
//
// 2022-09-28 06:53
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifferenceMinmaxWithStruct-12        68554383                16.94 ns/op
// BenchmarkMaxProductDifferenceMinmaxWithoutStruct-12     82901066                14.28 ns/op
// BenchmarkMaxProductDifferenceSort-12                    13251756                89.17 ns/op
// BenchmarkMaxProductDifferenceClone-12                     714526              1547 ns/op
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 4.781s
//
// 2022-09-27 23:24
//
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkMaxProductDifferenceMinmax-12          83021190                14.27 ns/op
// BenchmarkMaxProductDifferenceSort-12            12678550                93.19 ns/op
// BenchmarkMaxProductDifferenceClone-12             737697              1625 ns/op
// PASS
// ok      github.com/lloydlobo/leetcode/easy/1913-maximum-product-difference-between-two-pairs/go 4.698s
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
