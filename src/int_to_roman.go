package leetcode

func IntToRoman(num int) string {
	numThresholds := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanStr := ""
	thresholdI := 0
	for num > 0 {
		if num >= numThresholds[thresholdI] {
			romanStr += romanSymbols[thresholdI]
			num -= numThresholds[thresholdI]
		} else {
			thresholdI++
		}

	}
	return romanStr

}
