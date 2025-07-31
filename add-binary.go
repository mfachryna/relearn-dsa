package relearndsa

func addBinary(a string, b string) string {
	result := ""
	ap, bp, current := len(a)-1, len(b)-1, 0

	for ap >= 0 || bp >= 0 || current > 0 {
		val := byte(0)
		if ap >= 0 {
			val += a[ap] - '0'
			ap--
		}
		if bp >= 0 {
			val += b[bp] - '0'
			bp--
		}
		if current > 0 {
			val++
			current = 0
		}

		switch val {
		case 1:
			result = "1" + result
		case 2:
			result = "0" + result
			current++
		case 3:
			result = "1" + result
			current++
		default:
			result = "0" + result
		}
	}
	return result
}
