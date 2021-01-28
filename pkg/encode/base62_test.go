package encode

import "testing"

func Test_toBase62(t *testing.T) {
	tests := []struct {
		name string
		num  uint64
		want string
	}{
		{"zero", 0, ""},
		{"normal", 1056, "RC"},
		{"negative number", 18446744073709551615, "V8qRkBGKRiP"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBase62(tt.num); got != tt.want {
				t.Errorf("toBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}
