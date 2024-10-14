package leetcode

func majorityElement(nums []int) int {
	freqMap := map[int]int{}
	threshold := len(nums) / 2

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		freq, exists := freqMap[num]

		if freq >= threshold {
			//We can stop the computation early
			return num
		}
		if exists {
			freqMap[num]++
		} else {
			freqMap[num] = 1
		}
	}
	return -1
}
