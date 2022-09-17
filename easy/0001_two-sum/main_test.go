package main

// https://gobyexample.com/testing
// Writing tests can be repetitive, so it’s idiomatic to use a table-driven style, where test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic.
import (
	"fmt"
	"testing"
	// "regexp"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasics(t *testing.T) {
	output := IntMin(2, -2)
	if output != -2 {
		// t.Error* will report test failures but continue executing the test.
		t.Errorf("IntMin(2, -2) = %d; want -2", output)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)

		t.Run(testname, func(t *testing.T) {
			output := IntMin(tt.a, tt.b)
			if output != tt.want {
				t.Errorf("got %d, want %d", output, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		IntMin(1, 2)
	}
}

// t.Error* will report test failures but continue executing the test.
// t.Fatal* will report test failures and stop the test immediately.
/* func (*testing.T).Run(name string, f func(t *testing.T)) bool
   ─────────────────────────────────────────────────────────────────────────────────────────
   Run runs f as a subtest of t called name\. It runs f in a separate goroutine
   and blocks until f returns or calls t\.Parallel to become a parallel test\.
   Run reports whether f succeeded \(or at least did not fail before calling t\.Parallel\)\.
   Run may be called simultaneously from multiple goroutines, but all such calls
   must return before the outer test function for t returns\.
   [`(testing.T).Run` on pkg.go.dev](https://pkg.go.dev/testing?utm_source=gopls#T.Run) */
