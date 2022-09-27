package minmax

// Is it like counting? and tallying?
func MaxProductDifferenceMinMax(nums []int) int {
	max1, max2 := nums[0], nums[1]
	if max1 < max2 {
		max1, max2 = max2, max1
	}
	min1, min2 := nums[0], nums[1]
	if min1 > min2 {
		min1, min2 = min2, min1
	}

	size := len(nums)

	for i := 2; i < size; i++ {
		num := nums[i]

		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	return max1*max2 - min1*min2
}
