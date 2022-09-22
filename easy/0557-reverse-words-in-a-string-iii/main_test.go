// Given a string s, reverse the order of characters in each word,
// within a sentence while still preserving whitespace and initial word order.
//
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
package main

import (
	"log"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// func TestReverseWords(t *testing.InternalTest)  {
// reverseWords reverse the order of chars in each word.
func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	want := "s'teL ekat edoCteeL tsetnoc"
	got := reverseWords(s)
	// size := len(strings.Split(s, " "))
	if got != want {
		t.Errorf("reverseWords(%s) = %s; want: %s", s, got, want)
	}
}

// BenchmarkReverseWords()
func BenchmarkReverseWords(b *testing.B) {
	s := "Let's take LeetCode contest"
	for i := 0; i < b.N; i++ {
		reverseWords(s)
	}
}

// `sliceString` slices the string with seperator " ".
// and returns a slice of words.
//
// t.Errorf(cmp.Equal(got, want))
func TestSliceString(t *testing.T) {
	s := "Let's take LeetCode contest"
	size := len(strings.Split(s, " "))
	want := []string{"Let's", "take", "LeetCode", "contest"}
	got, gotSize := sliceString(s)
	log.Printf("sliceString(%s) = %s; want: %s; test: %v", s, got, want, cmp.Equal(got, want))
	if size != gotSize {
		t.Errorf("_, len(%s) = sliceString(s); want: %v; got: %v", s, size, gotSize)
	}
}

// `sliceString` slices the string with seperator " ".
// and returns a slice of words.
func BenchmarkSliceString(b *testing.B) {
	s := "Let's take LeetCode contest"
	for i := 0; i < b.N; i++ {
		sliceString(s)
	}
}

// ///////////////////////////////////////////////////////////////
// Benchmark Results.
// ///////////////////////////////////////////////////////////////

/*
# 2022-09-22 12:10
12:09  ➜  go test -bench=. -v
=== RUN   TestReverseWords
--- PASS: TestReverseWords (0.00s)
=== RUN   TestSliceString
2022/09/22 12:09:18 sliceString(Let's take LeetCode contest) = [Let's take LeetCode contest]; want: [Let's
 take LeetCode contest]; test: true
--- PASS: TestSliceString (0.00s)
goos: linux
goarch: amd64
pkg: github.com/lloydlobo/leetcode/easy/0557-reverse-words-in-a-string-iii
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkReverseWords
BenchmarkReverseWords-12          456525              2459 ns/op
BenchmarkSliceString
BenchmarkSliceString-12          3912076               308.2 ns/op
PASS
ok      github.com/lloydlobo/leetcode/easy/0557-reverse-words-in-a-string-iii   2.671s */

// ///////////////////////////////////////////////////////////////
// ARCHIVE.
// ///////////////////////////////////////////////////////////////
/*
    func splitLoop(s string) {
    	var sliceStrings []string
    	for i := 0; i < len(strings.Split(s, "")); i++ {
    		splitS := strings.Split(s, "")
    		// multi array?
    		sliceStrings = append(sliceStrings, splitS...)
    		fmt.Printf("splitS: %v\n", splitS)
    	}
    }

    func strings.Cut(s string, sep string) (before string, after string, found bool)
    sBefore, sAfter, isFound := strings.Cut(s, " ")

   	if isFound {
   		fmt.Printf("sBefore: %s | sAfter: %s |\n", sBefore, sAfter)
   	}
*/

/* var p = fmt.Println

func smain() {
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
} */
