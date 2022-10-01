// [[file:../../../README.org::0344-reverse-string/go][0344-reverse-string/go]]
package algo

// Write a function that reverses a string.
// - The input string is given as an array of characters s.
// - You must do this by modifying the input array in-place with O(1) extra memory.
// Constraints:
// 1 <= s.length <= 105
// s[i] is a printable ascii character.
//
// [104 101 108 108 111]
// [111 108 108 101 104]

// Runtime: 38 ms, faster than 79.95% of Go online submissions for Reverse String.
// Memory Usage: 6.8 MB, less than 41.35% of Go online submissions for Reverse String.
func ReverseString(s []byte) []byte {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return s
}

// Runtime: 54 ms
// Memory Usage: 6.9 MB
func ReverseStringStd(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseStringMid(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}

// 0344-reverse-string/go ends here
