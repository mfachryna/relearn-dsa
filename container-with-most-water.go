package relearndsa

func maxArea(height []int) int {
	firstP, lastP := 0, len(height)-1
	maxs := 0
	for firstP < lastP {
		hori := lastP - firstP
		maxs = max(maxs, min(height[firstP], height[lastP])*hori)

		if height[firstP] > height[lastP] {
			lastP--
		} else {
			firstP++
		}
	}
	return maxs
}
