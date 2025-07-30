package relearndsa

func mySqrt(x int) int {
	low, high := 0, x

	for low <= high {
		mid := (low + high) / 2
		power := mid * mid

		if power == x {
			return mid
		}

		if power < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low - 1
}
