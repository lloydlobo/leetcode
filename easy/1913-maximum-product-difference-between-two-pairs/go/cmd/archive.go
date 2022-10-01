package cmd

// //////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////
// ARCHIVE///////////////////////////////////////////////////
// //////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////

// var channel = make(chan (*MinMax))
// // https://gobyexample.com/channels
//
//	go func() {
//		channel <- InitMinMax(nums)
//	}()
//
// // m = <-channel
//
// func MaxProductDifference(nums []int) int {
// 	var (
// 		n        = len(nums)
// 		out      int
// 		hash     = make(map[int][]int, n)
// 		copyNums = nums
// 	)

// 	// Sort the array and cache the indexes in a map.
// 	sort.Slice(copyNums, func(i, j int) bool {
// 		isGreater := copyNums[i] > copyNums[j]

// 		if isGreater {
// 			keys := make([]int, 0, len(hash))

// 			for k := range hash {
// 				keys = append(keys, k)
// 				// fmt.Printf("keys: %v\n", keys)
// 			}
// 		}

// 		return isGreater
// 	})

// 	if n >= 4 && n <= 104 {
// 		for i := 0; i < n; i++ {
// 			num := nums[i]
// 			if num >= 1 && num <= 104 {
// 				out = product(nums, num, i, n)

// 			} else {

// 				return 0 // err
// 			} // end of if else nums

// 		} // end of for i loop

// 		return out
// 	} // end of if

// 	return 0 // err
// }

// // product()
// //
// // Input: nums = [5,6,2,7,4]
// // Output: 34
// // Explanation: We can choose indices 1 and 3 for the first pair (6, 7)
// // and indices 2 and 4 for the second pair (2, 4).
// // The product difference is (6 * 7) - (2 * 4) = 34.
// func product(nums []int, num, idx, size int) int {
// 	var (
// 		w    int
// 		x    int
// 		y    int
// 		z    int
// 		hash = make(map[int][]int, size)
// 	)
// 	cNums := nums
// 	// Sort the array and cache the indexes in a map.
// 	// sort.Ints(copyNums)
// 	sort.Slice(cNums, func(i, j int) bool {
// 		ci, cj := cNums[i], cNums[j]
// 		isGreater := ci > cj
// 		fmt.Printf("ci(%v): %v cj(%v): %v \n", i, ci, j, cj)
// 		return isGreater
// 	})
// 	fmt.Printf("copyNums: %v\n", cNums)

// 	keys := make([]int, 0, len(hash))
// 	for k := range hash {
// 		keys = append(keys, k)
// 		fmt.Printf("keys: %v\n", keys)
// 	}
// 	// Multiply two biggest numbers and two smallest numbers.

// 	fmt.Println(idx, w, x, y, z, hash)
// 	fmt.Println(idx, nums, cNums, hash)

// 	return num
// }
