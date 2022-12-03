package secondpuzzle

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		label    string
		input    [][]string
		expected int
	}{
		{"Should return 70", [][]string{{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, {"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}}, 70},
	}
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := Solve(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong output expected %v got %v\n", tc.expected, output)
			}
		})
	}
}
