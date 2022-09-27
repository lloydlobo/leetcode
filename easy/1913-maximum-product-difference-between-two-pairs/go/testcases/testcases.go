package testcases

import "log"

func GetTestcases() ([][]int, []int) {
	var (
		nums  [][]int
		wants []int
	)
	nums1, want1 := []int{5, 6, 2, 7, 4}, 34
	nums2, want2 := []int{4, 2, 5, 9, 7, 4, 8}, 64

	nums, wants = append(nums, nums1, nums2), append(wants, want1, want2)
	lenN, lenW := len(nums), len(wants)
	// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
	if lenN != lenW {
		log.SetPrefix("GetTestcases()")
		log.Default().Fatalf("lenN is not equal to lenW")
	}

	return nums, wants
}
