package testcase

import (
// "fmt"

// "github.com/fatih/color"
)

func ExecForLoop(f func([]int) bool, arrNums [][]int, arrWant []bool, n int) ([]bool, []bool) {
	var arrOutputGot []bool
	var arrOutPutWant []bool
	for i := 0; i < n; i++ {
		got := f(arrNums[i])
		arrOutputGot = append(arrOutputGot, got)
		want := arrWant[i]
		arrOutPutWant = append(arrOutPutWant, want)
		// color.Set(color.FgYellow, color.Bold)
		// fmt.Printf("➜ %2v | output: %11v | want: %11v\n\n", i, got, want)
		// color.Unset()
	}
	return arrOutputGot, arrOutPutWant
}

func GetTestcase() ([][]int, []bool) {
	arrNums := [][]int{{3, 4, 5, 1, 2}, {2, 1, 3, 4}, {1, 2, 3}}
	arrWants := []bool{true, false, true}
	return arrNums, arrWants
}
