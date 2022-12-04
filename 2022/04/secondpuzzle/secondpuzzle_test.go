package secondpuzzle

import "testing"

func TestSolve(t *testing.T) {
	input := [][][]int{{{2, 4}, {6, 8}}, {{2, 3}, {4, 5}}, {{5, 7}, {7, 9}}, {{2, 8}, {3, 7}}, {{6, 6}, {4, 6}}, {{2, 6}, {4, 8}}}
	output := Solve(input)
	if output != 4 {
		t.Fatalf("Should return 4 got %v\n", output)
	}
}
