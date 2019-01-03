package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			reply := sumIntsRedux(tt.input)
			if reply != tt.output {
				t.Errorf("Expected %v, got %v", tt.output, reply)
			}

		})
	}
}
