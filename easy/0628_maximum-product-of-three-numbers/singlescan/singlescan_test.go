package singlescan

import (
	"math"
	"testing"

	"github.com/lloydlobo/leetcode/testcases"
)

func TestMaximumProductSingleScan(t *testing.T) {
	arrNums, arrWant := testcases.GetMainArgs() // Get response from getArgs() test cases.
	for i := 0; i < len(arrNums); i++ {         // for loop to test each iteration in slice arrNums.
		got := MaximumProductSingleScan(arrNums[i]) // got is the output from the maximumProduct function.
		want := arrWant[i]                          // want is the expected result.
		if got != want {
			t.Errorf("TestMaximumProduct(): #%v | got: %v | want: %v |", i, got, want)
		}
	}
}

/*
# BenchmarkMaximumProductSingleScan

	2022-09-20 15:54
	goos: linux
	goarch: amd64
	pkg: github.com/lloydlobo/leetcode/singlescan
	cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	BenchmarkMaximumProductSingleScan-12            1000000000               0.0000513 ns/op
	BenchmarkRandInt-12                             570889324                2.071 ns/op
	2022-09-19 20:00
	BenchmarkMaximumProductSorting-12               18769503                61.97 ns/op
	BenchmarkMaximumProductSingleScan-12            126393331                9.520 ns/op
	2022-09-19 19:56
	BenchmarkMaximumProductSorting-12               20541787                57.89 ns/op
	BenchmarkMaximumProductSingleScan-12            123170702                9.974 ns/op
	2022-09-19 19:07
	BenchmarkMaximumProductSorting-12               29704512                38.49 ns/op
	BenchmarkMaximumProductSingleScan-12            145782132                8.209 ns/op
	BenchmarkRandInt-12                             564502472                2.088 ns/op
	PASS
	ok      github.com/lloydlobo/leetcode   4.614s
*/
func BenchmarkMaximumProductSingleScan(b *testing.B) {
	arrNums, _ := testcases.GetMainArgs() // Get response from getArgs() test cases.
	for i := 0; i < len(arrNums); i++ {   // for loop to test each iteration in slice arrNums.
		MaximumProductSingleScan(arrNums[i]) // got is the output from the maximumProduct function.
	}
}

/*
# BenchmarkRandInt

	=== RUN   TestMaximumProduct
	--- PASS: TestMaximumProduct (0.00s)
	goos: linux
	goarch: amd64
	pkg: github.com/lloydlobo/leetcode
	cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	BenchmarkRandInt
	BenchmarkRandInt-12     548181571                2.149 ns/op
	PASS
	ok      github.com/lloydlobo/leetcode   1.402s
*/
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Exp(float64(i))
	}
}
