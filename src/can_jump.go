package leetcode

func canJump(nums []int) bool {
	jump := 0
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		jump--
		jump = max(jump, nums[i])
		if jump == 0 {
			return false
		}
	}
	return true
}
