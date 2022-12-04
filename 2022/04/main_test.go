package main

import (
	"testing"
)

func TestConvertPairToSlice(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected []int
	}{
		{"Should return the slice", "2-4", []int{2, 4}},
		{"Should return the slice", "6-12", []int{6, 12}},
	}
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output, _ := convertPairToSlice(tc.input)
			if tc.expected[0] != output[0] || tc.expected[1] != output[1] {
				t.Fatalf("Wrong array returned expected %v got %v", tc.expected, output)
			}
		})
	}
}

func TestParseRow(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected [][]int
	}{
		{"Should return a length 2 slice", "2-4,6-12", [][]int{{2, 4}, {6, 12}}},
		{"Should return a length 2 slice", "2-8,6-12", [][]int{{2, 8}, {6, 12}}},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output, _ := parseRow(tc.input)
			if len(tc.expected) != len(output) {
				t.Fatalf("Wrong length returned\n")
			}
			if tc.expected[0][0] != output[0][0] || tc.expected[1][0] != output[1][0] {
				t.Fatalf("Wrong element returned\n")
			}
		})
	}
}

func TestParseData(t *testing.T) {
	cases := []struct {
		label    string
		input    []string
		expected [][][]int
	}{
		{"Should return a length 2 slice", []string{"2-4,6-12", "3-4,8-9"}, [][][]int{{{2, 4}, {6, 12}}, {{3, 4}, {8, 9}}}},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output, _ := parseData(tc.input)
			if len(output) != len(tc.expected) {
				t.Fatalf("Wrong number of elements returned expected %v got %v\n", len(tc.expected), len(output))
			}
			if output[0][0][0] != tc.expected[0][0][0] {
				t.Fatalf("Wrong elemnt returned\n")
			}
		})
	}
}
