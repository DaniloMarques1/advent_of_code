package firstpuzzle

func Solve(rows [][][]int) int {
	tot := 0
	for _, row := range rows {
		firstPair := row[0]
		secondPair := row[1]
		if pairContains(firstPair, secondPair) || pairContains(secondPair, firstPair) {
			tot += 1
		}
	}

	return tot
}

// will check if pair is contained inside pairToCompare
func pairContains(pair, pairToCompare []int) bool {
	return pairToCompare[0] <= pair[0] && pairToCompare[1] >= pair[1]
}
