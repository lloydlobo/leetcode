// [[file:../README.org::two-sum][two-sum]]
package main


import "fmt"

func twoSum(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i < len-1; i += 1 {
		for j := i + 1; j < len; j += 1 {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	fmt.Println("001_two-sum/main.go: main()")
	nums := []int{2, 7, 10, 15}
	target := 9
	output := twoSum(nums, target)
	fmt.Println("\nOutput:", output)
}
// two-sum ends here
