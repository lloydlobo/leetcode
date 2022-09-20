package main

import (
	"testing"

	testcase "github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated/testcase"
)

/*
# TestCheck()

	## 2022-09-20 19:24
	=== RUN   TestCheck
	--- PASS: TestCheck (0.00s)
	PASS
	ok      github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated    0.001s
*/
func TestCheck(t *testing.T) {
	arrNums, arrWants := testcase.GetTestcase()
	gots, wants := testcase.ExecForLoop(Check, arrNums, arrWants, len(arrNums))
	for i := 0; i < len(gots); i++ {
		if gots[i] != wants[i] {
			t.Errorf("TestCheck | got: %v; want: %v", gots[i], wants[i])
		}
	}
}

/*
BenchmarkCheck()

	2022-09-20 19:27
	 - === RUN   TestCheck
	 - --- PASS: TestCheck (0.00s)
	 - goos: linux
	 - goarch: amd64
	 - pkg: github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated
	 - cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	 - BenchmarkCheck
	 - BenchmarkCheck-12        9034135               129.1 ns/op
	 - PASS
	 - ok      github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated    1.305s
*/
func BenchmarkCheck(b *testing.B) {
	arrNums, arrWants := testcase.GetTestcase()
	for i := 0; i < b.N; i++ {
		testcase.ExecForLoop(Check, arrNums, arrWants, len(arrNums))
	}
}
