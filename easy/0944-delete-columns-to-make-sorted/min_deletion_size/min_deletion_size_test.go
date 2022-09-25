package min_deletion_size

import (
	"testing"
)

func TestMinDeletionSize(t *testing.T) {
	strs := []string{"cba", "daf", "ghi"}
	got, want := MinDeletionSize(strs), 1
	if 1 != 1 {
		t.Errorf("MinDeletionSize(%s) = %v; want: %v", strs, got, want)
	}
}

func BenchmarkMinDeletionSize(b *testing.B) {
	strs := []string{"cba", "daf", "ghi"}
	for i := 0; i < b.N; i++ {
		MinDeletionSize(strs)
	}
}

/*
   Test & Benchmark Log.

   # 2022-09-25 13:20
   === RUN   TestMinDeletionSize
   --- PASS: TestMinDeletionSize (0.00s)
   goos: linux
   goarch: amd64
   pkg: github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/min_deletion_size
   cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
   BenchmarkMinDeletionSize
   BenchmarkMinDeletionSize-12     1000000000           0.2340 ns/op
   PASS
   ok      github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/min_deletion_size   0.261s
*/
