// [[file:../README.org::palindrome-number][palindrome-number]]
package main


// works
import (
	"fmt"
	"strconv"
	// "strings"
)

/*
	 reverseStr() Function.
		  https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/
			type string string
			string is the set of all strings of 8\-bit bytes, conventionally but not necessarily representing UTF\-8\-encoded text\.
			A string may be empty, but not nil\. Values of string type are immutable\. [`string` on pkg.go.dev](https://pkg.go.dev/builtin?utm_source=gopls#string)
*/
func reverseStr(str string) (result string) {
	// Append the result at each for loop turn at the end of string.
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func isPalindrome(x int) bool {
	strX := strconv.FormatInt(int64(x), 10)
	strRevX := reverseStr(strX)
	return strX == strRevX
} // fmt.Println("strX:", strX, "strings", strRevX)

func printResult(int int) {
	result := isPalindrome(int)
	fmt.Println("Output:", int, "isPalindrome", result)
}

// ────────────────────────────────────────────────────────────────────────────
func main() {
	sliceIntNums := []int{121, 212, 123, 101}
	for _, intNum := range sliceIntNums {
		printResult(intNum)
	}
}
// palindrome-number ends here
