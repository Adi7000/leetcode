package leetcode

import "log"

func jump(nums []int) int {
	return performJump(nums, 1)
}

// Returns the number of jumps required to reach the end of nums
// Starts at one of the first positions in nums upto startUpto (non-inclusive)
func performJump(nums []int, startUpto int) int {
	//Base case is if end is already reachable
	if startUpto <= 0 {
		log.Fatal("startUpto must be greater than 0")
	}
	if startUpto >= len(nums) {
		return 0
	}
	maxJump := effectiveJump(nums[:startUpto])
	return 1 + performJump(nums[startUpto:], maxJump)
}

// Calculates the effective jump distance from any of the nums provided and returns it
// Jump distance is calculated from the end of nums
func effectiveJump(nums []int) int {
	lastPosition := len(nums) - 1
	effectiveJump := 0
	for i, jump := range nums {
		currentEffectiveJump := jump - (lastPosition - i)
		if currentEffectiveJump > effectiveJump {
			effectiveJump = currentEffectiveJump
		}
	}
	return effectiveJump
}
