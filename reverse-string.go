package relearndsa


func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
}

func reverseString2(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func reverseString3(s []byte) {
    i,l := 0, len(s)-1
    for i < l {
		s[i], s[l] = s[l], s[i]
        i++
        l--
    }
}


func init() {
    debug.SetMemoryLimit(1)
}
