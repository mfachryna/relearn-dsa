package relearndsa

func searchInsert(nums []int, target int) int {
	pointer := 0
	for key, value := range nums {
		if value == target {
			return key
		}
		if target >= value {
			pointer = key + 1
		}
	}

	return pointer
}
