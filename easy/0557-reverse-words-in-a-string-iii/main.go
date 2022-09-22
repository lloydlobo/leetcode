// 557. Reverse Words in a String III
//
// Given a string s, reverse the order of characters in each word
// within a sentence while still preserving whitespace and initial word order.
//
// Examples:
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Input: s = "God Ding"
// Output: "doG gniD"
//
// Constraints:
// - 1 <= s.length <= 5 * 104
// - s contains printable ASCII characters.
// - s does not contain any leading or trailing spaces.
// - There is at least one word in s.
// - All the words in s are separated by a single space.
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	// "image/color"
	// s "strings"
)

// `sliceString` slices the string with seperator " ".
// and returns a slice of words.
//
// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
// If sep is empty, Split splits after each UTF\-8 sequence\.
func sliceString(s string) ([]string, int) {
	size := len(strings.Split(s, " "))
	var arrS []string
	sSlice := strings.Split(s, " ")
	for i := 0; i < size; i++ {
		arrS = append(arrS, sSlice[i])
	}
	return arrS, size
}

// sliceChars()
//
// Slice each characters into a slice.
// Then reverse them in position.
func sliceChars(arrS []string, size int) []string {
	var split []string
	var arrChars [][]string
	var arrSize []int
	var output []string
	var arrOutput []string
	for i := 0; i < size; i++ {
		split = strings.Split(arrS[i], "")
		arrChars = append(arrChars, split)
		arrSize = append(arrSize, len(arrChars))
	}
	for i := 0; i < size; i++ {
		word := arrChars[i]
		for j := len(word) - 1; j >= 0; j-- {
			output = append(output, word[j])
		}
		arrOutput = append(arrOutput, strings.Join(output, ""))
		output = []string{""}
	}
	return arrOutput
}

// reverseWords reverse the order of chars in each word.
//
// Given a string s, reverse the order of characters in each word,
// within a sentence while still preserving whitespace and initial word order.
//
// Runtime: 49 ms.
// Memory Usage: 8.8 MB.
func reverseWords(s string) string {
	arrS, size := sliceString(s)
	output := sliceChars(arrS, size)
	return strings.Join(output, " ")
}

// Function mainPrintOutput sets log Flags and prefix,
// before printing the output to the console.
func mainPrintOutput(s string) {
	log.SetPrefix("main() ➜ ")
	log.SetFlags(0)
	log.Printf("output: %v\n", s)
}

// main runs and prints output for reverseWords().
func main() {
	tStart := time.Now()
	s := "Let's take LeetCode contest"
	output := reverseWords(s)
	tEnd := time.Now()
	tElapsed := tEnd.Sub(tStart)
	mainPrintOutput(output)
	fmt.Printf("\nCompleted in %v!\n", tElapsed)
}
