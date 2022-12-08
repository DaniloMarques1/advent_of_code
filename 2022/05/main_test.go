package main

import (
	"testing"
)

func TestParseData(t *testing.T) {
	cases := []struct {
		label    string
		input    []string
		expected [][]int
	}{
		{
			"Should return 3 plays",
			[]string{"move 1 from 7 to 5", "move 8 from 1 to 8", "move 1 from 1 to 9"},
			[][]int{{1, 7, 5}, {8, 1, 8}, {1, 1, 9}},
		},
		{
			"Should return 2 plays",
			[]string{"move 29 from 5 to 1", "move 2 from 6 to 3"},
			[][]int{{29, 5, 1}, {2, 6, 3}},
		},
		{
			"Should return 1 play",
			[]string{"move 1 from 2 to 1"},
			[][]int{{1, 2, 1}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(*testing.T) {
			output := parseData(tc.input)
			if len(output) != len(tc.expected) {
				t.Fatalf("Wrong output expected %v got %v\n", tc.expected, output)
			}

			if output[0][1] != tc.expected[0][1] {
				t.Fatalf("Wrong value returned. expected %v got %v\n", tc.expected[0][1], output[0][1])
			}
		})
	}
}

func TestIsNumericString(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected bool
	}{
		{"Should return false", "danilindo", false},
		{"Should return true", "23", true},
		{"Should return false", "123danlindo", false},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			_, output := isNumericString(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong output. expected %v got %v\n", tc.expected, output)
			}
		})
	}
}
