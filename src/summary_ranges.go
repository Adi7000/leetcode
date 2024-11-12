package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	ranges := []string{}
	var start *int = nil

	for i, num := range nums {
		if start == nil {
			start = &num
		} else {
			if num != nums[i-1]+1 {
				if nums[i-1] == *start {
					ranges = append(ranges, fmt.Sprintf("%d", *start))
				} else {
					ranges = append(ranges, fmt.Sprintf("%d->%d", *start, nums[i-1]))
				}
				start = &num
			}
		}
	}
	if start != nil {
		if nums[len(nums)-1] == *start {
			ranges = append(ranges, fmt.Sprintf("%d", *start))
		} else {
			ranges = append(ranges, fmt.Sprintf("%d->%d", *start, nums[len(nums)-1]))
		}
	}
	return ranges
}
