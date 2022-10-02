package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	algo "github.com/lloydlobo/leetcode/easy/1955-count-special-quadruplets/go/algo"
)

func TestRunMain(t *testing.T) {
	main()
}
func BenchmarkRunMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

//	func TestAlgo1(t *testing.T) {
//		got := algo.CountQuadruplets(Nums)
//		fmt.Printf("Nums: %v\n", Nums)
//		want := 1
//		if got != want {
//			t.Errorf("algo.CountQuadruplets(%v) = %v; want: %v!", Nums, got, want)
//		}
//	}

// - (0, 1, 2, 3): 1 + 1 + 1 == 3
// - (0, 1, 3, 4): 1 + 1 + 3 == 5
// - (0, 2, 3, 4): 1 + 1 + 3 == 5
// - (1, 2, 3, 4): 1 + 1 + 3 == 5
func TestAlgo0GetMap(t *testing.T) {
	got, want := len(algo.GetMap(Nums)), 5

	if got != want {
		t.Errorf("algo.GetMap(%v) = %v; want: %v!", Nums, got, want)
	}

}

func TestAlgo0SortLess(t *testing.T) {
	// got, want := algo.SortLess(Nums), []int{1, 2, 3, 6}
	got, want := algo.SortLess(Nums), []int{1, 1, 1, 3, 5}
	if !cmp.Equal(got, want) {
		t.Errorf("algo.SortLess(%v) = %v; want: %v!", Nums, got, want)
	}

}
func TestAlgo0GetLen(t *testing.T) {
	got, want := algo.GetLen(Nums), 5
	if got != want {
		t.Errorf("algo.GetLen(%v) = %v; want: %v!", Nums, got, want)
	}
}
func TestAlgo0Swap(t *testing.T) {
	n := len(Nums)
	a, b := 0, n-1
	algo.SwapAWithB(&Nums, a, b)
	gotA, gotB, wantA, wantB := Nums[a], Nums[b], 6, 1
	if gotA != wantA && gotB != wantB {
		t.Errorf("algo.SwapAWithB(%v), a: %v, b: %v; got: %v %v; want: %v %v ", Nums, a, b, gotA, gotB, wantA, wantB)
	}
}

func BenchmarkAlgo0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.CountQuadruplets(Nums)
	}
}
