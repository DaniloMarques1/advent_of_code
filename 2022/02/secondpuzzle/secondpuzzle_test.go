package secondpuzzle

import "testing"

func TestCalculateTotalScore(t *testing.T) {
	cases := []struct {
		label    string
		input    [][]string
		expected int
	}{
		{"Should return 12", [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}, 12},
		{"Should return 27", [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}, {"C", "Y"}, {"B", "Z"}}, 27},
		{"Should return 8", [][]string{{"A", "Z"}}, 8},
		{"Should return 4", [][]string{{"A", "Y"}}, 4},
		{"Should return 3", [][]string{{"A", "X"}}, 3},
		{"Should return 7", [][]string{{"C", "Z"}}, 7},
		{"Should return 7", [][]string{{"A", "Y"}, {"A", "Y"}, {"B", "X"}, {"A", "Y"}}, 13},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := calculateTotalScore(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong output got %v expected %v\n", output, tc.expected)
			}
		})
	}
}
