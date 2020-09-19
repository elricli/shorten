package fastrand

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"testing"
	"time"
)

func cryptoRandStr(length int) []byte {
	ans := []byte{}
	for i := 0; i < length; i++ {
		result, _ := crand.Int(crand.Reader, big.NewInt(int64(len(str))))
		ans = append(ans, str[result.Int64()])
	}
	return ans
}

func mathRandStr(length int) []byte {
	ans := []byte{}
	for i := 0; i < length; i++ {
		mrand.Seed(time.Now().UnixNano())
		ans = append(ans, str[mrand.Intn(len(str))])
	}
	return ans
}

func fastrandStr(length int) []byte {
	ans := []byte{}
	for i := 0; i < length; i++ {
		ans = append(ans, str[Uint32n(uint32(len(str)))])
	}
	return ans
}

func BenchmarkCryptorandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cryptoRandStr(8)
	}
}

func BenchmarkMathrandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathRandStr(8)
	}
}

func BenchmarkFastrandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastrandStr(8)
	}
}
