package helper

import "testing"

func TestRemoveIndex(t *testing.T) {
	var (
		strs = []string{"cba", "daf", "ghi"}
		i    = 1
	)
	got, want := RemoveIndex(strs, i), 1
	if 1 != 2/2 {
		t.Errorf("RemoveIndex(%s, %v) = %v; want: %v", strs, i, got, want)
	}
}

func BenchmarkRemoveIndex(b *testing.B) {
	var (
		strs  = []string{"cba", "daf", "ghi"}
		index = 1
	)

	for i := 0; i < b.N; i++ {
		RemoveIndex(strs, index)
	}
}

func TestRemoveIndexConcise(t *testing.T) {
	var (
		strs = []string{"cba", "daf", "ghi"}
		i    = 1
	)
	got, want := RemoveIndexConcise(strs, i), 1
	if 1 != 2/2 {
		t.Errorf("RemoveIndex(%s, %v) = %v; want: %v", strs, i, got, want)
	}
}

func BenchmarkRemoveIndexConcise(b *testing.B) {
	var (
		strs  = []string{"cba", "daf", "ghi"}
		index = 1
	)

	for i := 0; i < b.N; i++ {
		RemoveIndexConcise(strs, index)
	}
}

// https://stackoverflow.com/a/57213476
// func TestRemoveIndexConcise(t *testing.T) {
// 	var (
// 		strs = []string{"cba", "daf", "ghi"}
// 		i    = 1
// 	)
// 	got, want := RemoveIndexConcise(strs, i), 1
// 	if 1 != 2/2 {
// 		t.Errorf("RemoveIndex(%s, %v) = %v; want: %v", strs, i, got, want)
// 	}
// }

/*
$ go test -v -bench=.

# Log

* 2022-09-25 15:13
=== RUN   TestRemoveIndex
--- PASS: TestRemoveIndex (0.00s)
=== RUN   TestRemoveIndexConcise
--- PASS: TestRemoveIndexConcise (0.00s)
goos: linux
goarch: amd64
pkg: github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/helper
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkRemoveIndex
BenchmarkRemoveIndex-12                 25046972                46.53 ns/op
BenchmarkRemoveIndexConcise
BenchmarkRemoveIndexConcise-12          13708045                86.47 ns/op
PASS
ok      github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/helper    2.493s

* 2022-09-25 15:08
=== RUN   TestRemoveIndex
--- PASS: TestRemoveIndex (0.00s)
goos: linux
goarch: amd64
pkg: github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/helper
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkRemoveIndex
BenchmarkRemoveIndex-12         24574198                46.11 ns/op
PASS
ok      github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/helper    1.186s

*/
