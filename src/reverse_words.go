package leetcode

import (
	"strings"
)

func ReverseWords(s string) string {
	//Split on whitespace and trim
	splitStrings := strings.Fields(s)

	//Reverse
	n := len(splitStrings)
	for i := 0; i < len(splitStrings)/2; i++ {
		tempStr := splitStrings[i]
		splitStrings[i] = splitStrings[n-i-1]
		splitStrings[n-i-1] = tempStr
	}

	//Rejoin
	return strings.Join(splitStrings, " ")
}
