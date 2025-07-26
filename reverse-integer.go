package relearndsa

import "strconv"

func reverse(x int) int {
	isNeg := 1
	if x < 0 {
		isNeg = -1
	}

	num := []rune(strconv.Itoa(x * isNeg))

	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}

	x, err := strconv.Atoi(string(num))
	if err != nil {
		return 0
	}
	x = x * isNeg

	minI := -2147483648
	maxI := 2147483647
	if x > maxI || x < minI {
		return 0
	}
	return x
}
