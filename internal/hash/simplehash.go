package hash

// Uint32 return uint32 hash code.
func Uint32(key string) uint32 {
	hash := uint32(32)
	for i := 0; i < len(key); i++ {
		hash = hash<<5 + uint32(key[i])
	}
	return hash
}
