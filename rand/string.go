package rand

var (
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// String is return random string with length
func String(length uint32) string {
	b := []byte{}
	for i := uint32(0); i < length; i++ {
		b = append(b, str[Uint32n(uint32(len(str)))])
	}
	return string(b)
}
