package relearndsa

func isSubsequence(s string, t string) bool {
	point := 0

	if s == "" {
		return true
	}

	for key, _ := range t {
		if s[point] == t[key] {
			point++
		}
		if point > len(s)-1 {
			return true
		}
	}
	return false
}
