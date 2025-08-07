package relearndsa

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := make([]int, rowIndex+1)
	res[0], res[rowIndex] = 1, 1
	prev := 1

	for i := 1; i < rowIndex; i++ {
		temp := (prev * (rowIndex - i + 1)) / i
		res[i], prev = temp, temp
	}
	return res
}
