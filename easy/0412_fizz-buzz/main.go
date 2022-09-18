// [[file:../README.org::*main][main:1]]
package main

import (
	"fmt"
	"strconv"
)

// fizzBuzz function swaps any integer divisible by 3, 5, or 15.
//
// # Brute Force
// Given a number n, Run a loop from 1 to n
//
//	If a number is divisible by 15.Print FizzBuzz
//	Else If a number is divisible by 3.Print Fizz
//	Else a number is divisible by 3.Print Buzz
//
// # Constraints:
//   - `1 <= n <= 104`
//
// https://www.tutorialcup.com/interview/algorithm/fizz-buzz.html
// Time complexity = O(n) where n is the number till we have to print the values in the fizz buzz way.
// Space Complexity = O(n).
// https://leetcode.com/problems/fizz-buzz/submissions/
// Runtime: 10 ms.
// Memory Usage: 4.3 MB.
func fizzBuzz(n int) []string {
	var output []string

	for i := 1; i < n+1; i++ {
		output = append(output, fmt.Sprint(i))

		if (i)%15 == 0 {
			output[i-1] = "FizzBuzz"
		} else if (i)%3 == 0 {
			output[i-1] = "Fizz"
		} else if (i)%5 == 0 {
			output[i-1] = "Buzz"
		}
		// FIXME: Add an `else` statement to append `i` to `output` slice
		// when `i` is not divisible by 3,5,15.
	}
	return output
}

// FizzBuzzAppend()
//
// Time complexity = O(n) where n is the number till we have to print the values in the fizz buzz way.
// Space Complexity = O(n).
//
// Runtime: 7 ms.
// Memory Usage: 3.4 MB.
func FizzBuzzAppend(n int) []string {
	var arr = make([]string, 0, n) // The make built\-in function allocates and initializes an object of type slice, map, or chan (only).

	for i := 1; i < n+1; i++ {
		if (i)%15 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if (i)%3 == 0 {
			arr = append(arr, "Fizz")
		} else if (i)%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, strconv.Itoa(i)) // Itoa is equivalent to FormatInt(int64(i), 10).
		}
	}
	return arr
}

// main function to print the output to console.
func main() {
	n := 16
	fmt.Printf("fizzBuzz: %v\n", fizzBuzz(n))
	n = 16
	fmt.Printf("fizzBuzz: %v\n", FizzBuzzAppend(n))
}

//   conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?) stringintconv [11, 27]

// main:1 ends here
