package firstpuzzle

import (
	"testing"
)

func TestCalculateRockRound(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected int
	}{
		{"Should return 3", "A", 3},
		{"Should return 0", "B", 0},
		{"Should return 6", "C", 6},
		{"Should return 0", "P", 0},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := calculateRockRound(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong ouput, expected %v got %v\n", tc.expected, output)
			}
		})
	}
}

func TestCalculatePaperRound(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected int
	}{
		{"Should return 6", "A", 6},
		{"Should return 3", "B", 3},
		{"Should return 0", "C", 0},
		{"Should return 0", "P", 0},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := calculatePaperRound(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong ouput, expected %v got %v\n", tc.expected, output)
			}
		})
	}
}

func TestCalculateScissorsRound(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected int
	}{
		{"Should return 0", "A", 0},
		{"Should return 6", "B", 6},
		{"Should return 3", "C", 3},
		{"Should return 0", "P", 0},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := calculateScissorsRound(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong ouput, expected %v got %v\n", tc.expected, output)
			}
		})
	}
}

func TestCalculateMyRoundScore(t *testing.T) {
	cases := []struct {
		label       string
		firstInput  string
		secondInput string
		expected    int
	}{
		{"Should return 8", "A", "Y", 8},
		{"Should return 1", "B", "X", 1},
		{"Should return 6", "C", "Z", 6},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := calculateMyRoundScore(tc.firstInput, tc.secondInput)
			if output != tc.expected {
				t.Fatalf("Wrong output expected %v got %v\n", tc.expected, output)
			}

		})
	}
}

func TestCalculateTotalScore(t *testing.T) {
	input := [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}
	output := calculateTotalScore(input)
	expected := 15
	if output != expected {
		t.Fatalf("Wrong output got %v expected %v\n", output, expected)
	}
}
