package leetcode

func Rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 || len(nums) == 1 || len(nums) == k {
		return
	}

	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)

}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		swap_storage := nums[i]
		nums[i] = nums[len(nums)-1-i]
		nums[len(nums)-1-i] = swap_storage
	}
}
