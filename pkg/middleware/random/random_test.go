package random

import (
	"testing"
)

func TestRandom_String(t *testing.T) {
	type args struct {
		length   uint8
		charsets []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "gen with generate",
			args: args{
				length:   32,
				charsets: []string{},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.args.length, tt.args.charsets...)
			if len(got) != tt.want {
				t.Errorf("Random.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		length   uint8
		charsets []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.length, tt.args.charsets...); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
