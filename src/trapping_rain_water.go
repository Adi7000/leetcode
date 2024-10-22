package leetcode

func trap(height []int) int {
	// Preallocate arrays, default value is 0 which is important
	// for the left and right boundaries of the edges
	leftBoundary := make([]int, len(height))
	rightBoundary := make([]int, len(height))

	// Find all left boundaries and then right boundaries
	for i := 1; i < len(height); i++ {
		leftBoundary[i] = max(leftBoundary[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightBoundary[i] = max(rightBoundary[i+1], height[i+1])
	}

	//Now calcualte the total water at each section capping the minimum to 0
	totalWater := 0
	for i := 0; i < len(height); i++ {
		lowerBoundary := min(leftBoundary[i], rightBoundary[i])
		water := max(0, lowerBoundary-height[i])
		totalWater += water
	}
	return totalWater
}
