package day02

import "testing"

func TestCheckSum(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		option rune
		want   int
	}{
		{"1st", `5 1 9 5
7 5 3 3
2 4 6 8`, ' ', 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSum(tt.input, tt.option); got != tt.want {
				t.Errorf("CheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
