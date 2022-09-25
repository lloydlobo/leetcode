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
//
// Note: Take care of 2 edge cases here,
// the first word does not have a space before its first character.
// Similarly, the last word does not have a space after its last character.
//
// Algorithm
// By analyzing the above two key observations, we can derive the following algorithm,
//   - Find the starting and ending position of each word in the string.
//   - As a space character is a separator for each word,
//     we are finding the substrings having a space character,
//     before its first character and after its last character.
//   - For each identified word, reverse the characters of the word one by one.
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const space byte = 32

func reverseWordsByteSpace(s string) string {
	bts := []byte(s)
	first := 0
	for i := 1; i < len(bts); i++ {
		if bts[i] == space {
			for x, y := first, i-1; x < y; x++ {
				bts[x], bts[y] = bts[y], bts[x]
				y--
			}
			first = i + 1
		}
	}

	for x, y := first, len(bts)-1; x < y; x++ {
		bts[x], bts[y] = bts[y], bts[x]
		y--
	}
	return string(bts)
}

// C style.
func swap(s string, start, end int) {
	arrChar := strings.Split(s, " ")
	fmt.Printf("arrChar: %v\n", arrChar)
	for start < end {
		start++
		end--
		fmt.Printf("s: %v, e: %v", start, end)
	}
}

// C style.
func swapWord(s string) string {
	wStart, wEnd, _ := 0, 0, len(s)

	// for wEnd < full {
	swap(s, wStart, wEnd)
	// }

	return s
}

/* reverseword.c
void swap(char *s, int start, int end) {
  while (start < end) {
    char charTemp = s[start];
    s[start++] = s[end];
    s[end--] = charTemp;
  }
}


// Runtime: 10 ms.
// Memory:  6.9 MB.
char *reverseWords(char *s) {
  int wStart = 0, wEnd = 0, len = strlen(s);
  while (wEnd < len) {
    while (wEnd < len && s[wEnd] != ' ')
      wEnd++;
    swap(s, wStart, wEnd - 1);
    wStart = ++wEnd;
  }
  return s;
}
*/

// Runtime: 0 ms.
// Memory Usage: 7.8 MB.
func reverseWordsByte(s string) string {
	word := strings.Split(s, " ") // Split string by space.
	for k, v := range word {      // Iterate over each word.
		b := []byte(v)             // Convert each word chars to bytes array uint8.
		n := len(v)                // Length of current word.
		for i := 0; i < n/2; i++ { // For each word, we iterate n / 2 times.
			b[i], b[n-1-i] = b[n-1-i], b[i] // Swap the characters.
			word[k] = string(b)             // word[key] here is the reversed version of value.
		}
	}
	return strings.Join(word, " ") // build the final string
}

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
	var output string
	tStart := time.Now()
	s := "Let's take LeetCode contest"
	output = reverseWords(s)
	output = swapWord(s)
	tEnd := time.Now()

	tElapsed := tEnd.Sub(tStart)
	// _ = tEnd.Sub(tStart)
	mainPrintOutput(output)
	fmt.Printf("\nCompleted in %v!\n", tElapsed)
}
