package testcases

type TestcasesRectangle struct {
	r [][][]int
}

// func GetTestcases() (TestcasesRectangle, int) {
func GetTestcases() ([][][]int, int) {
	// tcr := TestcasesRectangle{}
	r1 := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	r2 := [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}

	// tcr.r = [][][]int{r1, r2}
	tcr := [][][]int{r1, r2}
	n := len(tcr)

	return tcr, n
}
