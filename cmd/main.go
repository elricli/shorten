package main

import (
	"fmt"
	"hash/crc32"
	"time"

	"github.com/drrrMikado/shorten/rand"
)

func main() {
	l := 1000000
	sArr := []string{}
	for i := 0; i < l; i++ {
		sArr = append(sArr, rand.String(10))
	}
	beginTime := time.Now()
	for i := 0; i < l; i++ {
		hashcode(sArr[i])
	}
	fmt.Println("hashcode time:", time.Now().Sub(beginTime))
	beginTime = time.Now()
	for i := 0; i < l; i++ {
		crc32.ChecksumIEEE([]byte(sArr[i]))
	}
	fmt.Println("crc32.ChecksumIEEE time:", time.Now().Sub(beginTime))
}

func hashcode(key string) uint32 {
	hash := uint32(32)
	for i := 0; i < len(key); i++ {
		hash = hash<<5 + uint32(key[i])
	}
	return hash
}
