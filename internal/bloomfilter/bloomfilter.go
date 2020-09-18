package bloomfilter

import (
	"math"
	"sync"

	"github.com/drrrMikado/shorten/internal/hash"
)

// BloomFilter .
type BloomFilter struct {
	mut              sync.Mutex
	bitArray         []uint32
	numBits          uint32
	numHashFunctions uint32
	fpp              float64
}

// New return a bloomFilter struct. fpp is probability of false positives
func New(expectedInsertions uint32, fpp float64) *BloomFilter {
	m := optimalNumOfBits(expectedInsertions, fpp)
	k := optimalNumOfHashFunctions(fpp)
	bf := &BloomFilter{
		bitArray:         make([]uint32, m),
		numBits:          m,
		numHashFunctions: k,
		fpp:              fpp,
		mut:              sync.Mutex{},
	}
	return bf
}

// Insert to bloomfilter's bitmap.
func (bf *BloomFilter) Insert(key []byte) {
	bf.mut.Lock()
	defer bf.mut.Unlock()
	h := uint32(0)
	for i := uint32(0); i < bf.numHashFunctions; i++ {
		h |= hash.Murmur3_32(key)
		bitPos := h % bf.numBits
		bf.bitArray[bitPos] |= 1
	}
}

// MightContain return true if key might contain.
func (bf *BloomFilter) MightContain(key []byte) bool {
	bf.mut.Lock()
	defer bf.mut.Unlock()
	h := uint32(0)
	for i := uint32(0); i < bf.numHashFunctions; i++ {
		h |= hash.Murmur3_32(key)
		bitPos := h % bf.numBits
		if bf.bitArray[bitPos] == 0 {
			return false
		}
	}
	return true
}

func optimalNumOfBits(expectedInsertions uint32, fpp float64) uint32 {
	// see https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
	return uint32(math.Round(-1 * float64(expectedInsertions) * math.Log(fpp) / (math.Ln2 * math.Ln2)))
}

func optimalNumOfHashFunctions(fpp float64) uint32 {
	// see https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
	return uint32(math.Round(-1 * math.Log2(fpp)))
}
