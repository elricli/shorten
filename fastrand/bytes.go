package fastrand

var (
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// String is return random string with length
func String(length uint32) string {
	return string(Bytes(length))
}

// Bytes return random byte slice.
func Bytes(length uint32) []byte {
	b := []byte{}
	for i := uint32(0); i < length; i++ {
		b = append(b, str[Uint32n(uint32(len(str)))])
	}
	return b
}
