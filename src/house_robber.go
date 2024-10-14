package leetcode

func rob(nums []int) int {
	//Initialize
	robbedHouseSliceToAmountMap := map[int]map[int]int{}
	for i := 0; i <= len(nums); i++ {
		robbedHouseSliceToAmountMap[i] = map[int]int{}
	}

	return robPartition(nums, 0, len(nums), robbedHouseSliceToAmountMap)
}

func robPartition(nums []int, low int, high int, robbedHouseSliceToAmountMap map[int]map[int]int) int {
	numHouses := high - low
	//Check if problem already solved
	amount, exists := robbedHouseSliceToAmountMap[low][high]
	if exists {
		return amount
		//Base cases
	} else if numHouses == 0 {
		amount = 0
	} else if numHouses == 1 {
		amount = nums[low]
	} else if numHouses == 2 {
		amount = max(nums[low], nums[low+1])
	} else {
		//Partition into 2 slices in the middle
		//Note middle always belongs to the right side
		middle := low + numHouses/2

		//We have to prevent adjacent houses from being robbed
		//In case 1 assume last house of p1 isn't robbed
		case1 := robPartition(nums, low, middle-1, robbedHouseSliceToAmountMap) + robPartition(nums, middle, high, robbedHouseSliceToAmountMap)
		//In case 2 assume first house of p2 isn't robbed
		case2 := robPartition(nums, low, middle, robbedHouseSliceToAmountMap) + robPartition(nums, middle+1, high, robbedHouseSliceToAmountMap)
		amount = max(case1, case2)
	}
	robbedHouseSliceToAmountMap[low][high] = amount
	return amount
}
