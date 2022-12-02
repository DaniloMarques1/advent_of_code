package secondpuzzle

func Solve(rounds [][]string) int {
	return calculateTotalScore(rounds)
}

func calculateTotalScore(rounds [][]string) int {
	totalScore := 0
	for _, round := range rounds {
		totalScore += calculateRound(round[0], round[1])
	}
	return totalScore
}

func calculateRound(first, second string) int {
	roundScore := 0
	switch first {
	case "A":
		roundScore = calculateRock(second)
	case "B":
		roundScore = calculatePaper(second)
	case "C":
		roundScore = calculateScissors(second)
	default:
		return 0
	}
	return roundScore
}

func calculateRock(input string) int {
	switch input {
	case "X":
		return 3
	case "Y":
		return 4
	case "Z":
		return 8
	default:
		return 0
	}
}

func calculatePaper(input string) int {
	switch input {
	case "X":
		return 1
	case "Y":
		return 5
	case "Z":
		return 9
	default:
		return 0
	}
}

func calculateScissors(input string) int {
	switch input {
	case "X":
		return 2
	case "Y":
		return 6
	case "Z":
		return 7
	default:
		return 0
	}
}
