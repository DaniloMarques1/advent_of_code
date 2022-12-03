package firstpuzzle

import "testing"

func TestGetCompartmentFromRuckSack(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected []string
	}{
		{"Should return the correct ruckstack", "vJrwpWtwJgWrhcsFMMfFFhFp", []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}},
		{"Should return the correct ruckstack", "PmmdzqPrVvPwwTWBwg", []string{"PmmdzqPrV", "vPwwTWBwg"}},
		{"Should return the correct ruckstack", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", []string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}},
		{"Should return empty string", "", []string{"", ""}},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			firstOutput, secondOutput := getCompartmentFromRuckSack(tc.input)
			if firstOutput != tc.expected[0] || secondOutput != tc.expected[1] {
				t.Fatalf("Wrong output, expected %v and %v got %v and %v\n", tc.expected[0], tc.expected[1], firstOutput, secondOutput)
			}
		})
	}
}

func TestBuildMapFromCompartment(t *testing.T) {
	cases := []struct {
		label    string
		input    string
		expected map[string]int
	}{
		{"Should return a map where all chars have 1 value", "aabcdefgabcdefg", map[string]int{"a": 1, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1, "g": 1}},
		{"Should return a map where all chars have 1 value", "aeiou", map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}},
		{"Should return an empty map", "", map[string]int{}},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := buildMapFromCompartment(tc.input)
			if len(output) != len(tc.expected) {
				t.Fatalf("Number of keys returned is wrong\n")
			}
			if output["a"] != tc.expected["a"] {
				t.Fatalf("Wrong output\n")
			}
		})
	}
}

func TestCountOccurrencesFromSecondCompartment(t *testing.T) {
	cases := []struct {
		label       string
		firstInput  map[string]int
		secondInput map[string]int
		expected    []string
	}{
		{"Should increse the letter a to two", map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}, map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}, []string{"a", "e", "i", "o", "u"}},
		{"Should increse the letter a to two", map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}, map[string]int{"d": 1, "b": 1, "c": 2}, []string{}},
	}
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := countOccurrences(tc.firstInput, tc.secondInput)
			if len(tc.expected) != len(output) {
				t.Fatalf("Wrong number of elements returned expected %v got %v", len(tc.expected), len(output))
			}
		})
	}
}

func TestSumRepeatedItemsPriority(t *testing.T) {
	cases := []struct {
		label    string
		input    []string
		expected int
	}{
		{"Should return 10", []string{"a", "i"}, 10},
		{"Should return 10", []string{"p", "L"}, 54},
		{"Should return 10", []string{"P", "v"}, 64},
		{"Should return 10", []string{}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := sumRepeatedItemsPriority(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong value returned expected %v got %v\n", tc.expected, output)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	cases := []struct {
		label    string
		input    []string
		expected int
	}{
		{"Should return 157", []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}, 157},
		{"Should return 16", []string{"vJrwpWtwJgWrhcsFMMfFFhFp"}, 16},
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
