package hash

func murmur_32_scramble(k uint32) uint32 {
	k *= 0xcc9e2d51
	k = (k << 15) | (k >> 17)
	k *= 0x1b873593
	return k
}

func Murmur3_32(key []uint32, seed uint32) uint32 {
	h := seed
	len := uint32(len(key))
	var k uint32
	for i := len >> 2; i > 0; i-- {

	}
	/* Read the rest. */
	k = 0
	for i := len & 3; i > 0; i-- {
		k <<= 8
		k |= key[i-1]
	}
	// A swap is *not* necessary here because the preceding loop already
	// places the low bytes in the low places according to whatever endianness
	// we use. Swaps only apply when the memory is copied in a chunk.
	h ^= murmur_32_scramble(k)
	/* Finalize. */
	h ^= len
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
