package encode

func ToBase62(num uint64) string {
	encoded := ""
	for num > 0 {
		r := num % 62
		num /= 62
		encoded = string(characterSet62[r]) + encoded
	}
	return encoded
}
