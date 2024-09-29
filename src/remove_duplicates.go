package leetcode

func removeDuplicates(nums []int) int {
	lastUniqueNumberIndex := 0
	numberOfRemovedDuplicates := 0
	length := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastUniqueNumberIndex] {
			lastUniqueNumberIndex = i
			length++
		} else if i-lastUniqueNumberIndex >= 2 {
			numberOfRemovedDuplicates++
		} else {
			length++
		}
		nums[length] = nums[i]
	}
	return length + 1
}
