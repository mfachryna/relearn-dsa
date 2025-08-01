package relearndsa

func multiply(num1 string, num2 string) string {
	result := make([]byte, len(num1)+len(num2))

	if num1 == "0" || num2 == "0" {
		return "0"
	}
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			result[i+j+1] += ((num1[i] - '0') * (num2[j] - '0'))
			if result[i+j+1] >= 10 {
				result[i+j] += (result[i+j+1] / 10)
				result[i+j+1] %= 10
			}
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}
	for i := 0; i < len(result); i++ {
		result[i] += '0'
	}
	return string(result)
}
