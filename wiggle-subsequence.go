package relearndsa

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func wiggleMaxLength(nums []int) int {
	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			up = max(up, down+1)
			continue
		}
		if nums[i] > nums[i-1] {
			down = max(down, up+1)

		}
	}
	return max(up, down)
}
