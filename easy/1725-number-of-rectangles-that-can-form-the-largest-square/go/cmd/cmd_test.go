package cmd

import (
	"testing"

	"github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/go/testcases"
	"github.com/stretchr/testify/assert"
)

func TestCountGoodRectangles(t *testing.T) {
	var got []int
	want := []int{3, 3}
	arr, n := testcases.GetTestcases()

	for i := 0; i < n; i++ {
		got = append(got, CountGoodRectangles(arr[i]))
		if got[i] != want[i] {
			t.Errorf("Execute() = %v; want: %v", got[i], want[i])
		}
	}
}

func TestExecute(t *testing.T) {
	got := Execute()
	want := []int{3, 3}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Execute() = %v; want: %v", got[i], want[i])
		}
	}
}

// 2022-09-27 07:07
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/go/cmd
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkExecute-12                      3581187               337.5 ns/op
// BenchmarkCountGoodRectangles-12          4332554               275.0 ns/op
// PASS
// ok      github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/go/cmd    3.026s
func BenchmarkCountGoodRectangles(b *testing.B) {
	r1 := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	r2 := [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}
	r := [][][]int{r1, r2}

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(r); j++ {
			CountGoodRectangles(r[j])
		}
	}
}

// 2022-09-27 07:07
// BenchmarkExecute()
// goos: linux
// goarch: amd64
// pkg: github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/go/cmd
// cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
// BenchmarkExecute-12      3445204               342.8 ns/op
// PASS
// ok      github.com/lloydlobo/leetcode/easy/0944-delete-columns-to-make-sorted/go/cmd    1.538s
func BenchmarkExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Execute()
	}
}

func TestSomething(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")
	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	// assert.Nil(t, object)
	// assert for not nil (good when you expect something)
	/* if assert.NotNil(t, object) {
	   // now we know that object isn't nil, we are safe to make
	   // further assertions without causing any errors
	   assert.Equal(t, "Something", object.Value)
	 } */
}
