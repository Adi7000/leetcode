package leetcode

func removeElement(nums []int, val int) int {
	swap := func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	totalEqual := 0
	for i := 0; i < len(nums); i++ {
		if nums[i-totalEqual] == val {
			totalEqual++
			swap(i-totalEqual+1, len(nums)-totalEqual)
		}
	}
	return len(nums) - totalEqual
}
