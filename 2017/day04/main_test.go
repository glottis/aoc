package main

import (
	"testing"
)

func Test_checkPhrase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1st", args{"aa bb cc dd ee"}, true},
		{"2nd", args{"aa bb cc dd aa"}, false},
		{"3rd", args{"aa bb cc dd aaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPhrase(tt.args.s); got != tt.want {
				t.Errorf("checkPhrase() = %v, want %v, input: %v", got, tt.want, tt.args.s)
			}
		})
	}
}

func Test_checkPhraseRedux(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1st", args{"abcde fghij"}, true},
		{"2nd", args{"abcde xyz ecdab"}, false},
		{"3rd", args{"oiii ioii iioi iiio"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPhraseRedux(tt.args.s); got != tt.want {
				t.Errorf("checkPhraseRedux() = %v, want %v", got, tt.want)
			}
		})
	}
}
