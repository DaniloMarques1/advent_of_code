package firstpuzzle

import "testing"

func TestPop(t *testing.T) {
	cases := []struct {
		label     string
		input     []string
		expected1 []string
		expected2 string
	}{
		{"Should remove the last element", []string{"A", "B", "C"}, []string{"A", "B"}, "C"},
		{"Should remove the last element", []string{}, []string{}, ""},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output1, output2 := pop(tc.input)
			if len(output1) != len(tc.expected1) {
				t.Fatalf("Wrong output expected %v got %v\n", tc.expected1, output1)
			}

			if len(output1) > 0 {
				if output1[0] != tc.expected1[0] {
					t.Fatalf("Wrong output expected %v got %v\n", tc.expected1[0], output1[0])
				}

				if output2 != tc.expected2 {
					t.Fatalf("Wrong output expected %v got %v\n", tc.expected2, output2)
				}
			}
		})
	}
}

func TestRearrange(t *testing.T) {
	cases := []struct {
		label    string
		input1   map[int][]string
		input2   []int
		expected map[int][]string
	}{
		{
			"First step",
			map[int][]string{1: {"Z", "N"}, 2: {"M", "C", "D"}, 3: {"P"}},
			[]int{1, 2, 1},
			map[int][]string{1: {"Z", "N", "D"}, 2: {"M", "C"}, 3: {"P"}},
		},
		{
			"First step",
			map[int][]string{1: {"Z", "N", "D"}, 2: {"M", "C"}, 3: {"P"}},
			[]int{3, 1, 3},
			map[int][]string{1: {}, 2: {"M", "C"}, 3: {"P", "D", "N", "Z"}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := rearrange(tc.input1, tc.input2)
			if len(output) != len(tc.expected) {
				t.Fatalf("Wrong output returnd expected %v got %v\n", tc.expected, output)
			}

			for key, crates := range output {
				for idx, crate := range crates {
					if tc.expected[key][idx] != crate {
						t.Fatalf("Wrong output returnd expected %v got %v\n", tc.expected, output)
					}
				}
			}
		})
	}
}

func TestPlay(t *testing.T) {
	cases := []struct {
		label    string
		input1   map[int][]string
		input2   [][]int
		expected map[int][]string
	}{
		{
			"Should rearrange the crates",
			map[int][]string{1: {"Z", "N"}, 2: {"M", "C", "D"}, 3: {"P"}},
			[][]int{{1, 2, 1}, {3, 1, 3}, {2, 2, 1}, {1, 1, 2}},
			map[int][]string{1: {"C"}, 2: {"M"}, 3: {"P", "D", "N", "Z"}},
		},
		{
			"Should rearrange the crates",
			map[int][]string{1: {"A"}, 2: {"B"}, 3: {"C"}},
			[][]int{{1, 1, 2}, {1, 3, 2}},
			map[int][]string{1: {}, 2: {"B", "A", "C"}, 3: {}},
		},
		{
			"Should rearrange the crates",
			map[int][]string{1: {"A"}, 2: {"B"}, 3: {"C"}},
			[][]int{{1, 1, 2}, {1, 3, 2}, {1, 2, 1}},
			map[int][]string{1: {"C"}, 2: {"B", "A"}, 3: {}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := Play(tc.input1, tc.input2)
			for key, crates := range output {
				for idx, crate := range crates {
					if tc.expected[key][idx] != crate {
						t.Fatalf("Wrong output returnd expected %v got %v\n", tc.expected, output)
					}
				}
			}
		})
	}
}

func TestGetTopCrates(t *testing.T) {
	cases := []struct {
		label    string
		input    map[int][]string
		expected string
	}{
		{"Should return CMZ", map[int][]string{1: {"C"}, 2: {"M"}, 3: {"P", "D", "N", "Z"}}, "CMZ"},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			output := getTopCrates(tc.input)
			if output != tc.expected {
				t.Fatalf("Wrong output returned expected %v got %v\n", tc.expected, output)
			}
		})
	}
}
