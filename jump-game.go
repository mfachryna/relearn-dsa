package relearndsa

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := target - 1; i >= 0; i-- {
		if target <= i+nums[i] {
			target = i
		}
	}
	return target == 0
}
