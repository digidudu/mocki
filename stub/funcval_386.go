package stub

func funcValue(to uintptr) []byte {
	return []byte{
		0xBA,
		byte(to),
		byte(to >> 8),
		byte(to >> 16),
		byte(to >> 24), 
		0xFF, 0x22,     
	}
}