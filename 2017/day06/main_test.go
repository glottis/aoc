package main

import "testing"

func Test_countCycles(t *testing.T) {
	tests := []struct {
		name string
		ints []int
		want int
	}{
		{"first", []int{0, 2, 7, 0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCyclesRedux(tt.ints); got != tt.want {
				t.Errorf("countCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}
