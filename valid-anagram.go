package relearndsa

func isAnagram(s string, t string) bool {
	countRune := make([]int, 128)
	if len(s) != len(t) {
		return false
	}

	for _, value := range s {
		countRune[value-'a']++
	}

	for _, value := range t {
		if countRune[value-'a'] == 0 {
			return false
		}
		countRune[value-'a']--
	}

	return true
}
