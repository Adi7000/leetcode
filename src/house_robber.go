package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	//Partition into 2 slices in the middle
	//Note middle always belongs to the right side
	middle := (len(nums)) / 2

	//We have to prevent adjacent houses from being robbed
	//In case 1 assume last house of p1 isn't robbed
	case1 := rob(nums[:middle-1]) + rob(nums[middle:])
	//In case 2 assume first house of p2 isn't robbed
	case2 := rob(nums[:middle]) + rob(nums[middle+1:])
	return max(case1, case2)
}
