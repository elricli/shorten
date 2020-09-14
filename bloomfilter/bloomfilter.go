package bloomfilter

import (
	"math"

	"github.com/drrrMikado/shorten/hash"
)

// BloomFilter .
type BloomFilter struct {
	bitmap      []uint32
	bitmapSize  uint32
	numHashFunc uint32
	fpp         float64
}

// New return a bloomFilter struct. fpp is probability of false positives
func New(bitmapSize int, fpp float64) *BloomFilter {
	// see https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
	m := uint32(-1 * float64(bitmapSize) * math.Log(fpp) / (math.Ln2 * math.Ln2))
	k := uint32(-1 * math.Log2(fpp))
	bf := &BloomFilter{
		bitmap:      make([]uint32, m),
		bitmapSize:  m,
		numHashFunc: k,
		fpp:         fpp,
	}
	return bf
}

// Insert to bloomfilter's bitmap.
func (bf *BloomFilter) Insert(key []byte) {
	h := uint32(0)
	for i := uint32(0); i < bf.numHashFunc; i++ {
		h |= hash.Murmur3_32(key)
		bitPos := h % bf.bitmapSize
		bf.bitmap[bitPos] |= 1
	}
}

// MightContain return true if key might contain.
func (bf *BloomFilter) MightContain(key []byte) bool {
	h := uint32(0)
	for i := uint32(0); i < bf.numHashFunc; i++ {
		h |= hash.Murmur3_32(key)
		bitPos := h % bf.bitmapSize
		if bf.bitmap[bitPos] == 0 {
			return false
		}
	}
	return true
}
