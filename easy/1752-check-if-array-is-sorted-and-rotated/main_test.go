package main

import (
	"testing"

	testcase "github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated/testcase"
)

/*
# TestCheck()

	2022-09-20 19:24
	=== RUN   TestCheck
	--- PASS: TestCheck (0.00s)
	PASS
	ok      github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated    0.001s
*/
func TestCheck(t *testing.T) {
	tc := &testcase.Testcase{}
	tc.ArrNums, tc.ArrWants = testcase.GetTestcase()
	tc.Length = len(tc.ArrNums)
	gots, wants := testcase.ExecForLoop(Check, tc)
	for i := 0; i < len(gots); i++ {
		if gots[i] != wants[i] {
			t.Errorf("TestCheck | got: %v; want: %v", gots[i], wants[i])
		}
	}
}

// === RUN   TestCheckMiss
// --- PASS: TestCheckMiss (0.00s)
func TestCheckMiss(t *testing.T) {
	tc := &testcase.Testcase{}
	tc.ArrNums, tc.ArrWants = testcase.GetTestcase()
	tc.Length = len(tc.ArrNums)
	gots, wants := testcase.ExecForLoop(CheckMiss, tc)
	for i := 0; i < len(gots); i++ {
		if gots[i] != wants[i] {
			t.Errorf("TestCheck | got: %v; want: %v", gots[i], wants[i])
		}
	}
}

/*
BenchmarkCheck()

	2022-09-20 19:27
	 - goos: linux
	 - goarch: amd64
	 - pkg: github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated
	 - cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	 - BenchmarkCheck-12        9034135               129.1 ns/op
	 - PASS
	 - ok      github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated    1.305s
*/
func BenchmarkCheck(b *testing.B) {
	tc := &testcase.Testcase{}
	tc.ArrNums, tc.ArrWants = testcase.GetTestcase()
	tc.Length = len(tc.ArrNums)
	for i := 0; i < b.N; i++ {
		testcase.ExecForLoop(Check, tc)
	}
}

/*
# BenchmarkCheckMiss()

	goos: linux
	goarch: amd64
	pkg: github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated
	cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz

	2022-09-21 09:19
	BenchmarkCheck-12                4925289               241.3 ns/op
	BenchmarkCheckMiss-12            6939687               168.1 ns/op
	PASS
	ok      github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated    2.782s


	2022-09-21 05:12  ➜  go test -v -bench=.
	BenchmarkCheck-12                4953817               237.5 ns/op
	BenchmarkCheckMiss-12            6815260               169.5 ns/op
	PASS
	ok      github.com/lloydlobo/leetcode/easy/1752-check-if-array-is-sorted-and-rotated    2.761s
*/
func BenchmarkCheckMiss(b *testing.B) {
	tc := &testcase.Testcase{}
	tc.ArrNums, tc.ArrWants = testcase.GetTestcase()
	tc.Length = len(tc.ArrNums)
	for i := 0; i < b.N; i++ {
		testcase.ExecForLoop(CheckMiss, tc)
	}
}
