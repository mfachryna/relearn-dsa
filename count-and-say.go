package relearndsa

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return converter(countAndSay(n - 1))
}

func converter(say string) string {
	result := strings.Builder{}
	counter := 1
	i := 1
	for ; i < len(say); i++ {
		if say[i] == say[i-1] {
			counter++
		} else {
			result.WriteString(strconv.Itoa(counter))
			result.WriteByte(say[i-1])
			counter = 1
		}
	}
	result.WriteString(strconv.Itoa(counter))
	result.WriteByte(say[i-1])
	return result.String()
}
