package day01

import "testing"

func Test_praseList(t *testing.T) {

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
				t.Errorf("praseList() = name %v, got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
