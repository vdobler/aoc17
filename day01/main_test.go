package main

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	} {
		t.Run(tc.in, func(t *testing.T) {
			if got := answer(tc.in); got != tc.want {
				t.Errorf("answer(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

func TestAnswer2(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	} {
		t.Run(tc.in, func(t *testing.T) {
			if got := answer2(tc.in); got != tc.want {
				t.Errorf("answer2(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
