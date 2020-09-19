package fastrand

import (
	"sync"
	"time"
)

// Uint32 returns pseudorandom uint32.
//
// It is safe calling this function from concurrent goroutines.
func Uint32() uint32 {
	v := prngPool.Get()
	if v == nil {
		v = &PRNG{}
	}
	r := v.(*PRNG)
	x := r.uint32()
	prngPool.Put(r)
	return x
}

var prngPool sync.Pool

// Uint32n returns pseudorandom uint32 in the range [0..maxN).
//
// It is safe calling this function from concurrent goroutines.
func Uint32n(maxN uint32) uint32 {
	x := Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

// PRNG is a pseudorandom number generator.
//
// It is unsafe to call PRNG methods from concurrent goroutines.
type PRNG struct {
	x uint32
}

// uint32 returns pseudorandom uint32.
//
// It is unsafe to call this method from concurrent goroutines.
func (r *PRNG) uint32() uint32 {
	for r.x == 0 {
		r.x = getRandomUint32()
	}

	// See https://en.wikipedia.org/wiki/Xorshift
	x := r.x
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.x = x
	return x
}

func getRandomUint32() uint32 {
	x := time.Now().UnixNano()
	return uint32((x >> 32) ^ x)
}
