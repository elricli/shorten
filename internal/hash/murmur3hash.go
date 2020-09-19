package hash

func murmur32Scramble(k uint) uint {
	k *= 0xcc9e2d51
	k = (k << 15) | (k >> 17)
	k *= 0x1b873593
	return k
}

// Murmur3_32 murmur3hash .
// https://en.wikipedia.org/wiki/MurmurHash
func Murmur3_32(key []byte, seed uint) uint {
	h := seed
	var k uint
	for i := len(key) >> 2; i > 0; i-- {
		k = uint(key[0]) | uint(key[1])<<8 | uint(key[2])<<16 | uint(key[3])<<24
		key = key[4:]
		h ^= murmur32Scramble(k)
		h = (h << 13) | (h >> 19)
		h = h*5 + 0xe6546b64
	}
	/* Read the rest. */
	k = 0
	for i := len(key) & 3; i > 0; i-- {
		k <<= 8
		k |= uint(key[i-1])
	}
	// A swap is *not* necessary here because the preceding loop already
	// places the low bytes in the low places according to whatever endianness
	// we use. Swaps only apply when the memory is copied in a chunk.
	h ^= murmur32Scramble(k)
	/* Finalize. */
	h ^= uint(len(key))
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
