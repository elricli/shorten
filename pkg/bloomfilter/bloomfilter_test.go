package bloomfilter

import (
	"os"
	"testing"
	"time"

	"github.com/drrrMikado/shorten/pkg/fastrand"
)

var (
	bf *BloomFilter
)

func TestMain(m *testing.M) {
	var err error
	bf, err = New(10000000, 0.00001, 0x1)
	if err != nil {
		os.Exit(1)
	}
	for i := 0; i < 9990000; i++ {
		str := fastrand.String(7)
		if !bf.MightContain([]byte(str)) {
			bf.Insert([]byte(str))
		} else {
			i--
		}
	}
	os.Exit(m.Run())
}

func BenchmarkBloomfilter(b *testing.B) {
	bf2 := &BloomFilter{}
	*bf2 = *bf
	b.ResetTimer()
	initNum := bf2.NumUsed
	s := time.Now()
	for i := 0; i < b.N; i++ {
		for {
			str := fastrand.String(7)
			if bf2.MightContain([]byte(str)) {
				continue
			}
			bf2.Insert([]byte(str))
			break
		}
	}
	b.Log("init num:", initNum, "insert", b.N, ",spend time:", time.Now().Sub(s), "nums:", bf2.NumUsed)
}
