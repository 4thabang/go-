package main

import (
	"testing"
)

func TestSumMapReduce(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{"happy path", []int{1, 2, 3, 4, 5}, 15},
		{"edge case #1", []int{}, 0},
		{"edge case #2", []int{2, 4, 5}, 11},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumMapReduce(tc.input)
			if got != tc.want {
				t.Errorf("got: %d, want: %d", got, tc.want)
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
}
