package relearndsa

import "math"

func titleToNumber(columnTitle string) int {
	number := 0
	for i := 0; i < len(columnTitle); i++ {
		number += (int(math.Pow(26, float64(len(columnTitle)-i-1))) * (int(columnTitle[i]-'A') + 1))
	}
	return number
}
