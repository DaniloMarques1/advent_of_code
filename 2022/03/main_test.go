package main

import "testing"

func TestParseDataForSecondPuzzle(t *testing.T) {
	cases := []struct {
		label    string
		input    []string
		expected [][]string
	}{
		{"Should return a length 2", []string{"abcdef", "abcdef", "abcdef", "abcdef", "abcdef", "abcdef"}, [][]string{{"abcdef", "abcdef", "abcdef"}, {"abcdef", "abcdef", "abcdef"}}},
		{"Should return a length 5", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"}, [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}}},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := parseDataForSecondPuzzle(tc.input)
			if len(output) != len(tc.expected) {
				t.Fatalf("Wrong number of elements returned expected %v got %v\n", len(tc.expected), len(output))
			}
		})
	}

}
