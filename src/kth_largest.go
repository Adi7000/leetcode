package leetcode

func findKthLargest(nums []int, k int) int {

	max, min := findMinMax(nums)

	frequencyNums := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		frequencyNums[nums[i]-min]++
	}
	for i := len(frequencyNums) - 1; i >= 0; i-- {
		if frequencyNums[i] < k {
			k -= frequencyNums[i]
		} else {
			return min + i
		}
	}
	return 0
}

func findMinMax(nums []int) (int, int) {
	max := nums[0]
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	return max, min
}
