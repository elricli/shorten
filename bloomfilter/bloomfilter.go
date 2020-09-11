package bloomfilter

import "math"

// BloomFilter .
type BloomFilter struct {
	bitmap []uint32
	m      uint32
	k      uint32
}

// New return a bloomFilter struct. fpp is probability of false positives
func New(bitmapSize int, fpp float64) *BloomFilter {
	m := uint32(-1 * float64(bitmapSize) * math.Log(fpp) / (math.Ln2 * math.Ln2))
	k := uint32(-1 * math.Log2(fpp))
	bf := &BloomFilter{
		bitmap: make([]uint32, m),
		m:      m,
		k:      k,
	}
	return bf
}

func hashcode(key string) uint32 {
	hash := uint32(32)
	for i := 0; i < len(key); i++ {
		hash = 31*hash + uint32(key[i])
	}
	return hash
}
