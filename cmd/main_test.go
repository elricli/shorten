package main

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"sync"
	"testing"
	"time"

	"github.com/drrrMikado/shorten/rand"
)

var (
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randomStr1(length int) []byte {
	ans := []byte{}
	for i := 0; i < length; i++ {
		result, _ := crand.Int(crand.Reader, big.NewInt(int64(len(str))))
		ans = append(ans, str[result.Int64()])
	}
	return ans
}

func randomStr2(length int) []byte {
	ans := []byte{}
	for i := 0; i < length; i++ {
		mrand.Seed(time.Now().UnixNano())
		ans = append(ans, str[mrand.Intn(len(str))])
	}
	return ans
}

func randomStr3(length int) []byte {
	ans := []byte{}
	for i := 0; i < length; i++ {
		ans = append(ans, str[rand.Uint32n(uint32(len(str)))])
	}
	return ans
}

func BenchmarkRandomStr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomStr1(8)
	}
}

func BenchmarkRandomStr2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomStr2(8)
	}
}

func BenchmarkRandomStr3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomStr3(8)
	}
}

func BenchmarkCheckStrExist(b *testing.B) {
	m := map[string]bool{}
	l := sync.Mutex{}
	c := 0
	b.SetParallelism(20)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for {
				s := rand.String(10)
				l.Lock()
				if _, ok := m[s]; !ok {
					m[s] = true
					l.Unlock()
					break
				} else {
					c++
					l.Unlock()
				}
			}
		}
	})
	if c > 0 {
		b.Log("conflict count:", c)
	}

}

func bloomFilter() {

}
