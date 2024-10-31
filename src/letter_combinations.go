package leetcode

func letterCombinations(digits string) []string {
	digitToLetters := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if len(digits) == 0 {
		return []string{}
	}
	return rec(digits, digitToLetters)

}

func rec(digits string, digitToLetters map[string]string) []string {
	if len(digits) == 0 {
		return []string{""}
	}
	subwords := rec(digits[1:], digitToLetters)
	numSubWords := len(subwords)
	newLetters := digitToLetters[string(digits[0])]
	for i := 1; i < len(newLetters); i++ {
		subwords = append(subwords, subwords[:numSubWords]...)
	}
	for i, letter := range newLetters {
		letter := string(letter)
		for j := 0; j < numSubWords; j++ {
			subwords[i*numSubWords+j] = letter + subwords[i*numSubWords+j]
		}
	}
	return subwords
}
