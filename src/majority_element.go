package leetcode

func majorityElement(nums []int) int {
	result, count := 0, 0
	for _, num := range nums {
		switch {
		case num == result:
			count++
		case count == 0:
			result = num
			count = 1
		default:
			count--
		}
	}
	return result
}
