package relearndsa

func convertToTitle(columnNumber int) string {
	var res []rune
	for columnNumber > 0 {
		columnNumber--
		res = append([]rune{'A' + rune(columnNumber%26)}, res...)
		columnNumber /= 26
	}
	return string(res)
}
