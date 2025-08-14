package relearndsa

func reverseBits(n int) int {
	var result int = 0
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= n & 1
		n >>= 1
	}
	return result

}
