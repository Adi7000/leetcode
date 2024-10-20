package leetcode

func trap(height []int) int {
	leftBoundary := make([]int, len(height))
	rightBoundary := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		boundary := height[i]
		for j := 0; j < len(height); j++ {
			if j < i {
				rightBoundary[j] = max(rightBoundary[j], boundary)
			} else {
				leftBoundary[j] = max(leftBoundary[j], boundary)
			}
		}
	}
	totalWater := 0
	for i := 0; i < len(height); i++ {
		lowerBoundary := min(leftBoundary[i], rightBoundary[i])
		water := max(0, lowerBoundary-height[i])
		totalWater += water

	}
	return totalWater
}
