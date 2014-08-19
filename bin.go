package bin

func decode16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func encode16(b []byte, i int, v uint16) {
	b[i] = byte(v)
	b[i+1] = byte(v >> 8)
}

func decode32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func encode32(b []byte, i int, v uint32) {
	b[i] = byte(v)
	b[i+1] = byte(v >> 8)
	b[i+2] = byte(v >> 16)
	b[i+3] = byte(v >> 24)
}

func decode64(b []byte) uint64 {
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func encode64(b []byte, i int, v uint64) {
	b[i] = byte(v)
	b[i+1] = byte(v >> 8)
	b[i+2] = byte(v >> 16)
	b[i+3] = byte(v >> 24)
	b[i+4] = byte(v >> 32)
	b[i+5] = byte(v >> 40)
	b[i+6] = byte(v >> 48)
	b[i+7] = byte(v >> 56)
}
