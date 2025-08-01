package relearndsa

func climbStairs(n int) int {
	first, sec := 0, 1
	for ; n > 0; n-- {
		first, sec = sec, first+sec
	}
	return sec
}
