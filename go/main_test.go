package main

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	tests := map[string]struct {
		a    string
		b    string
		want string
	}{
		"1": {"11", "1", "100"},
		"2": {"1010", "1011", "10101"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := addBinary(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("got %s instead of %s", got, tc.want)
			}
		})
	}
}

func Test_isRobotBounded(t *testing.T) {
	tests := map[string]struct {
		instructions string
		want         bool
	}{
		"1": {"GGLLGG", true},
		"2": {"GG", false},
		"3": {"GL", true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isRobotBounded(tc.instructions)

			if got != tc.want {
				t.Errorf("got %t instead of %t", got, tc.want)

			}
		})
	}
}

func Test_nextGreatestLetter(t *testing.T) {
	tests := map[string]struct {
		letters []byte
		target  byte
		want    byte
	}{
		"1": {[]byte{'c', 'f', 'j'}, 'a', 'c'},
		"2": {[]byte{'c', 'f', 'j'}, 'c', 'f'},
		"3": {[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := nextGreatestLetter(tc.letters, tc.target)

			if got != tc.want {
				t.Errorf("got %c instead of %c", got, tc.want)

			}
		})
	}
}
