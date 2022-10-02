package algo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// cmp.Equal() | Equal reports whether x and y are equal .

func inputTests() []byte {
	return []byte{104, 101, 108, 108, 111}
}

func TestReverseString(t *testing.T) {
	want := []byte{111, 108, 108, 101, 104}
	out := ReverseString(inputTests())
	if !cmp.Equal(out, want) {
		t.Errorf("ReverseString() = %v, want: %v", out, want)
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString(inputTests())
	}
}
func TestReverseStringStd(t *testing.T) {
	want := []byte{111, 108, 108, 101, 104}
	out := ReverseStringStd(inputTests())
	if !cmp.Equal(out, want) {
		t.Errorf("ReverseString() = %v, want: %v", out, want)
	}
}

func BenchmarkReverseStringStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseStringStd(inputTests())
	}
}
func TestReverseStringMid(t *testing.T) {
	want := []byte{111, 108, 108, 101, 104}
	out := ReverseStringMid(inputTests())
	if !cmp.Equal(out, want) {
		t.Errorf("ReverseString() = %v, want: %v", out, want)
	}
}

func BenchmarkReverseStringMid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseStringMid(inputTests())
	}
}

/*
$ go test -bench=.

2022-10-02 13:53

goos: linux
goarch: amd64
pkg: github.com/lloydlobo/leetcode/easy/0344-reverse-string/go/algo
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkReverseString-12               428896934                2.745 ns/op
BenchmarkReverseStringStd-12            434242290                2.642 ns/op
BenchmarkReverseStringMid-12            598683981                1.986 ns/op
PASS
ok      github.com/lloydlobo/leetcode/easy/0344-reverse-string/go/algo  4.285s


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
