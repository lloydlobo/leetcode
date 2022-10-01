// s = ["h","e","l","l","o"] Output: ["o","l","l","e","h"]
// s = ["H","a","n","n","a","h"] Output: ["h","a","n","n","a","H"]
package testcases

import "fmt"

func Input() [][]byte {
	var (
		iBytes []byte
		out    [][]byte
	)

	inputs := [][]string{
		{"h", "e", "l", "l", "o"},
		{"H", "a", "n", "n", "a", "h"},
	}

	for i := 0; i < len(inputs); i++ {
		s := inputs[i]
		for j := 0; j < len(s); j++ {
			b := byte(s[j][0])
			iBytes = append(iBytes, b)
		}
	}

	return iBytes
}

func Output() (wants [][]string) {

	wants = [][]string{
		{"o", "l", "l", "e", "h"},
		{"h", "a", "n", "n", "a", "H"},
	}
	return
}
