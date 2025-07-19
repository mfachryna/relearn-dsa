package relearndsa

func twoSum(nums []int, target int) []int {
	data := map[int]int{}
	for i := 0; i < len(nums); i++ {
		value, ok := data[nums[i]]
		if ok {
			return []int{value, i}
		}
		data[target-nums[i]] = i
	}
	return []int{}
}
