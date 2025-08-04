package relearndsa

func jump(nums []int) int {
	maxReach, jumps, curr, length := 0, 0, 0, len(nums)
	for i := 0; i < length-1; i++ {
		if maxReach < i+nums[i] {
			maxReach = i + nums[i]
		}

		if i == curr {
			jumps++
			curr = maxReach
		}
	}
	return jumps
}
