package day01

import (
	"testing"
)

func TestParseList(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int64
	}{
		{"test: 1111", []byte("1111"), 4},
		{"test: 1122", []byte("1122"), 3},
		{"test: 1234", []byte("1234"), 0},
		{"test: 91212129", []byte("91212129"), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseList(tt.input); got != tt.want {
				t.Errorf("ParseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseListRedux(t *testing.T) {

	tests := []struct {
		name  string
		input []byte
		want  int64
	}{
		{"test: 1212", []byte("1212"), 6},
		{"test: 1221", []byte("1221"), 0},
		{"test: 123425", []byte("123425"), 4},
		{"test: 123123", []byte("123123"), 12},
		{"test: 12131415", []byte("12131415"), 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseListRedux(tt.input); got != tt.want {
				t.Errorf("ParseListRedux() = %v, want %v", got, tt.want)
			}
		})
	}
}
