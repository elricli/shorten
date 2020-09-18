package validator

import "testing"

func TestIsURL(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{str: ""}, false},
		{"", args{str: "foobar.com"}, true},
		{"", args{str: "http://www.-foobar.com"}, false},
		{"", args{str: "http://www.foo----bar.com"}, false},
		{"", args{str: "/abs/test/path"}, false},
		{"", args{str: "http://foo&bar.org"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.str); got != tt.want {
				t.Errorf("IsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
