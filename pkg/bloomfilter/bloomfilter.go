package bloomfilter

import (
	"errors"
	"math"
	"sync"

	"github.com/drrrMikado/shorten/pkg/hash"
)

// BloomFilter .
type BloomFilter struct {
	mut              *sync.Mutex
	bitArray         []byte
	numBits          uint
	numHashFunctions uint
	fpp              float64
	hashSeed         uint
	NumUsed          uint
	digest           hash.Murmur3
}

// New return a bloomFilter struct. fpp is probability of false positives
func New(expectedInsertions uint, fpp float64, hashSeed uint) (*BloomFilter, error) {
	if expectedInsertions > math.MaxUint64 {
		return nil, errors.New("expected insertions too large")
	}
	m := optimalNumOfBits(expectedInsertions, fpp)
	k := optimalNumOfHashFunctions(fpp)
	return &BloomFilter{
		mut:              &sync.Mutex{},
		bitArray:         make([]byte, m),
		numBits:          m,
		numHashFunctions: k,
		fpp:              fpp,
		hashSeed:         hashSeed,
		digest:           hash.NewMurmur3(hashSeed),
	}, nil
}

// Insert to bloom filter's bit array.
func (bf *BloomFilter) Insert(key []byte) {
	bf.mut.Lock()
	defer bf.mut.Unlock()
	var h uint
	for i := uint(0); i < bf.numHashFunctions; i++ {
		h |= bf.digest.Sum32(key)
		bitPos := h % bf.numBits
		bf.bitArray[bitPos] |= 1
	}
	bf.NumUsed++
}

// MightContain return true if key might contain.
func (bf *BloomFilter) MightContain(key []byte) bool {
	bf.mut.Lock()
	defer bf.mut.Unlock()
	var h uint
	for i := uint(0); i < bf.numHashFunctions; i++ {
		h |= bf.digest.Sum32(key)
		bitPos := h % bf.numBits
		if bf.bitArray[bitPos] == 0 {
			return false
		}
	}
	return true
}

func optimalNumOfBits(expectedInsertions uint, fpp float64) uint {
	// see https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
	return uint(math.Round(-1 * float64(expectedInsertions) * math.Log(fpp) / (math.Ln2 * math.Ln2)))
}

func optimalNumOfHashFunctions(fpp float64) uint {
	// see https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
	return uint(math.Round(-1 * math.Log2(fpp)))
}
