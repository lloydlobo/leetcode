package testcase

type Testcase struct {
	ArrNums  [][]int
	ArrWants []bool
	Length   int
}

// func ExecForLoop(f func([]int) bool, arrNums [][]int, arrWant []bool, n int) ([]bool, []bool) {
func ExecForLoop(f func([]int) bool, t *Testcase) ([]bool, []bool) {
	var arrGot []bool
	var arrWant []bool
	t.ArrNums, t.ArrWants = GetTestcase()
	t.Length = len(t.ArrNums)
	for i := 0; i < t.Length; i++ {
		arrGot = append(arrGot, f(t.ArrNums[i]))
		arrWant = append(arrWant, t.ArrWants[i])
	}
	return arrGot, arrWant
}

func GetTestcase() ([][]int, []bool) {
	arrNums := [][]int{{3, 4, 5, 1, 2}, {2, 1, 3, 4}, {1, 2, 3}}
	arrWants := []bool{true, false, true}
	return arrNums, arrWants
}
