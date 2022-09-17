// Package testing provides support for automated testing of Go packages.
// It is intended to be used in concert with the "go test" command, which automates
// execution of any function of the form
//
//	func TestXxx(*testing.T)
//
// where Xxx does not start with a lowercase letter. The function name
// serves to identify the test routine.
//
// Within these functions, use the Error, Fail or related methods to signal failure.
//
// To write a new test suite, create a file whose name ends _test.go that
// contains the TestXxx functions as described here. Put the file in the same
// package as the one being tested. The file will be excluded from regular
// package builds but will be included when the "go test" command is run.
// For more detail, run "go help test" and "go help testflag".
package main

import (
	"math"
	"math/rand"
	"testing"
	// "crypto/rand"
)

func getTestSearchArgs() ([]int, int, int) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	want := 4

	return nums, target, want
}

// TestSearch is for testing binary search algorithm
// to find index of required target int.
//
// PASS
// ok      binary-search   0.001s
//
// --- FAIL: TestSearch (0.00s)
// main_test.go:36: Search([-1 0 3 5 8 12], 9) = -1; want 4
// FAIL
// exit status 1
// FAIL    binary-search   0.001s
func TestSearch(t *testing.T) {
	nums, target, want := getTestSearchArgs()

	got := Search(nums, target)
	if got != want {
		t.Errorf("Search(%d, %v) = %v; want %v", nums, target, got, want)
	}
}

// To run all benchmarks, use -bench=., as shown below:
//
//	go test -bench=.
//
// goos: linux
// goarch: amd64
// pkg: binary-search
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkSearch-12              426310108                2.815 ns/op
// BenchmarkRandInt-12             95099158                12.29 ns/op
// BenchmarkPrimeNumbers-12          471997              2384 ns/op
// BenchmarkBigLen-12                     1        1810085852 ns/op
// PASS
// ok      binary-search   5.634s
func BenchmarkSearch(b *testing.B) {
	nums, target, _ := getTestSearchArgs()

	for i := 0; i < b.N; i++ {
		Search(nums, target)
	}
}

// # Benchmarks Results
//
// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// goos: linux
// goarch: amd64
// pkg: binary-search
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkRandInt-12             88516600                12.22 ns/op
// BenchmarkPrimeNumbers-12          470682              2391 ns/op
// BenchmarkBigLen-12                     1        1866591573 ns/op
// PASS
// ok      binary-search   4.119s

// # Benchmarks
//
// Functions of the form
//
//	func BenchmarkXxx(*testing.B)
//
// are considered benchmarks, and are executed by the "go test" command when
// its -bench flag is provided. Benchmarks are run sequentially.
//
// For a description of the testing flags, see
// https://golang.org/cmd/go/#hdr-Testing_flags.
//
// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// SIMPLE BENCHMARKING SUITE
// https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func primeNumbers(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

// To run all benchmarks, use -bench=., as shown below:
//
// $ go test -bench=.
//
// b.N specifies the number of iterations; the value is not fixed, but dynamically allocated, ensuring that the benchmark runs for at least one second by default.
func BenchmarkPrimeNumbers(b *testing.B) {
	num := 121
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}

// The benchmark function must run the target code b.N times.
// During benchmark execution, b.N is adjusted until the benchmark function lasts
// long enough to be timed reliably. The output
//
//	BenchmarkRandInt-8   	68453040	        17.8 ns/op
//
// means that the loop ran 68453040 times at a speed of 17.8 ns per loop.
//
// If a benchmark needs some expensive setup before running, the timer
// may be reset:
func BenchmarkBigLen(b *testing.B) {
	// big := NewBig(40)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// big.Len()
		// fmt.Printf("big: %v\n", big)
		NewBig(43)
	}
}

func NewBig(n int) int {
	output := fib(n)
	return output
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	output := fib(n-1) + fib(n-2)
	return output
}

// SIMPLE TESTING SUITE
//
// Abs returns the absolute value of x.
func Abs(i int) int {
	output := math.Abs(float64(i))

	return int(output)
}

// TestAbs for Warmup testing
// T is a type passed to Test functions to manage test state and support formatted test logs.
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
