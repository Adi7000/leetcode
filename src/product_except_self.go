package leetcode

func productExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}

	runningProductForward := make([]int, len(nums))
	runningProductBackward := make([]int, len(nums))
	productExceptSelfNums := make([]int, len(nums))

	//Calculate running product forward like Ex. {1, 2, 3, 4} -> {1, 2, 6, 24}
	runningProductForward[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		runningProductForward[i] = runningProductForward[i-1] * nums[i]
	}

	//Calculate running product backward like  Ex. {1, 2, 3, 4} -> {24, 24, 12, 4}
	runningProductBackward[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		runningProductBackward[i] = runningProductBackward[i+1] * nums[i]
	}

	//The two neighbouring runningProducts are the prefix and suffix
	productExceptSelfNums[0] = runningProductBackward[1]
	productExceptSelfNums[len(nums)-1] = runningProductForward[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		productExceptSelfNums[i] = runningProductForward[i-1] * runningProductBackward[i+1]
	}
	return productExceptSelfNums
}

func productExceptSelfMemOptimized(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}
	productExceptSelfNums := make([]int, len(nums))
	runningProductForward := 1

	//Calculate running product backward like  Ex. {1, 2, 3, 4} -> {24, 24, 12, 4}
	productExceptSelfNums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		productExceptSelfNums[i] = productExceptSelfNums[i+1] * nums[i]
	}

	//Calculate running product forward like Ex. {1, 2, 3, 4} -> {1, 2, 6, 24}
	//No need to store RPF so simultaneously set the product except self value
	for i := 0; i < len(nums)-1; i++ {
		productExceptSelfNums[i] = runningProductForward * productExceptSelfNums[i+1]
		runningProductForward *= nums[i]
	}

	productExceptSelfNums[len(nums)-1] = runningProductForward

	return productExceptSelfNums
}
