package main

import "testing"

func Test_calcSteps(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"first", 1, 0},
		{"second", 12, 3},
		{"third", 23, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcSteps(tt.input)
			if err != nil {
				t.Errorf("calcSteps() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("calcSteps() = %v, want %v, input %v", got, tt.want, tt.input)
			}
		})
	}
}
