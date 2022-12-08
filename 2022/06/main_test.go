package main

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		label    string
		input1   rune
		input2   []rune
		expected bool
	}{
		{"Should return true", 'c', []rune{'c', 'b', 'a'}, true},
		{"Should return false", 'e', []rune{'c', 'b', 'a'}, false},
		{"Should return false", 'e', []rune{}, false},
		{"Should return true", 'e', []rune{'e'}, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := contains(tc.input1, tc.input2)
			if output != tc.expected {
				t.Fatalf("Wrong output expected %v got %v\n", tc.expected, output)
			}
		})
	}
}

func TestFindMarker(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected int
	}{
		{"Should return 5", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"Should return 6", "nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"Should return 10", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"Should return 11", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := findMarker(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong output. expected %v got %v\n", tc.expected, output)
			}
		})
	}
}

func TestFindStartOfMessage(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected int
	}{
		{"Should return 19", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"Should return 23", "bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"Should return 23", "nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"Should return 29", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"Should return 26", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := findStartOfMessage(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong output. expected %v got %v\n", tc.expected, output)
			}
		})
	}
}
