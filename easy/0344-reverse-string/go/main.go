package main

import (
	"github.com/lloydlobo/leetcode/easy/0344-reverse-string/go/algo"
)

func main() {
	_ = algo.ReverseString(input1())
	_ = algo.ReverseString(input2())

	_ = algo.ReverseStringStd(input1())
	_ = algo.ReverseStringStd(input2())

	_ = algo.ReverseStringMid(input1())
	_ = algo.ReverseStringMid(input2())
}

// %!(EXTRA []uint8=[72 97 110 110 97 104])
func input1() []byte {
	return []byte{104, 101, 108, 108, 111}
}
func input2() []byte {
	return []byte{72, 97, 110, 110, 97, 104}
}
