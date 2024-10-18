package leetcode

func lengthOfLastWord(s string) int {
	wordStartedFlag := false
	wordLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			wordLength++
			if !wordStartedFlag {
				wordStartedFlag = true
			}
		} else if wordStartedFlag {
			break
		}
	}
	return wordLength
}
