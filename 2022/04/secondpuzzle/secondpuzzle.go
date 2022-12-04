package secondpuzzle

func Solve(rows [][][]int) int {
	tot := 0

	for _, row := range rows {
		firstPair := row[0]
		secondPair := row[1]
		if doesOverlap(firstPair, secondPair) && doesOverlap(secondPair, firstPair) {
			tot += 1
		}
	}

	return tot
}

// 5, 8 - 3, 4

func doesOverlap(pair1, pair2 []int) bool {
	return pair1[1] >= pair2[0]
}
