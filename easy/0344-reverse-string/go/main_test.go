package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lloydlobo/leetcode/easy/0344-reverse-string/go/algo"
)

// cmp.Equal() | Equal reports whether x and y are equal .

func TestRunMain(t *testing.T) {
	main()
}
func BenchmarkRunMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func inputTests() []byte {
	return []byte{104, 101, 108, 108, 111}
}

func TestAlgoReverseString(t *testing.T) {
	want := []byte{111, 108, 108, 101, 104}
	out := algo.ReverseString(inputTests())
	if !cmp.Equal(out, want) {
		t.Errorf("ReverseString() = %v, want: %v", out, want)
	}
}

func BenchmarkAlgoReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.ReverseString(inputTests())
	}
}
func TestAlgoReverseStringStd(t *testing.T) {
	want := []byte{111, 108, 108, 101, 104}
	out := algo.ReverseStringStd(inputTests())
	if !cmp.Equal(out, want) {
		t.Errorf("ReverseString() = %v, want: %v", out, want)
	}
}

func BenchmarkAlgoReverseStringStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.ReverseStringStd(inputTests())
	}
}
func TestAlgoReverseStringMid(t *testing.T) {
	want := []byte{111, 108, 108, 101, 104}
	out := algo.ReverseStringMid(inputTests())
	if !cmp.Equal(out, want) {
		t.Errorf("ReverseString() = %v, want: %v", out, want)
	}
}

func BenchmarkAlgoReverseStringMid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.ReverseStringMid(inputTests())
	}
}

/*

2022-10-01 12:44

goos: linux
goarch: amd64
pkg: github.com/lloydlobo/leetcode/easy/0344-reverse-string/go
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkRunMain-12                     163435347             7.207 ns/op
BenchmarkAlgoReverseString-12           424779840             2.806 ns/op
BenchmarkAlgoReverseStringStd-12        427962609             2.799 ns/op
BenchmarkAlgoReverseStringMid-12        586916742             2.031 ns/op
PASS
ok      github.com/lloydlobo/leetcode/easy/0344-reverse-string/go    6.285s


2022-10-01 12:43

goos: linux
goarch: amd64
pkg: github.com/lloydlobo/leetcode/easy/0344-reverse-string/go
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkRunMain-12                     166226386                7.224 ns/op
BenchmarkAlgoReverseString-12           430190511                2.785 ns/op
PASS
ok      github.com/lloydlobo/leetcode/easy/0344-reverse-string/go       3.416s

*/
