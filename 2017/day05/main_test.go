package main

import "testing"

func Test_traverseInstruction(t *testing.T) {
	type args struct {
		intS []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1st", args{[]int{0, 3, 0, 1, -3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverseInstruction(tt.args.intS); got != tt.want {
				t.Errorf("traverseInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverseInstructionRedux(t *testing.T) {
	type args struct {
		intS []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1st", args{[]int{0, 3, 0, 1, -3}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := traverseInstructionRedux(tt.args.intS); got != tt.want {
				t.Errorf("traverseInstructionRedux() = %v, want %v", got, tt.want)
			}
		})
	}
}
