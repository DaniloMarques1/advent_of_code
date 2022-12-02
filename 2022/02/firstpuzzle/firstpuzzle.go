package firstpuzzle

func Solve(rounds [][]string) int {
	return calculateTotalScore(rounds)
}

func calculateTotalScore(rounds [][]string) int {
	totScore := 0
	for _, round := range rounds {
		totScore += calculateMyRoundScore(round[0], round[1])
	}
	return totScore
}

/*
first column we have A, B, C meaning rock paper and scissors
second column we have X, Y, Z meaning rock paper and scissors
rock beats scissors, paper beats rock and scissors beats paper
*/
func calculateMyRoundScore(other, mine string) int {
	roundScore := 0
	switch mine {
	case "X":
		roundScore += 1
		roundScore += calculateRockRound(other)
	case "Y":
		roundScore += 2
		roundScore += calculatePaperRound(other)
	case "Z":
		roundScore += 3
		roundScore += calculateScissorsRound(other)
	default:
		break
	}

	return roundScore
}

func calculateRockRound(hand string) int {
	switch hand {
	case "A":
		return 3
	case "B":
		return 0
	case "C":
		return 6
	default:
		return 0
	}
}

func calculateScissorsRound(hand string) int {
	switch hand {
	case "A":
		return 0
	case "B":
		return 6
	case "C":
		return 3
	default:
		return 0
	}
}

func calculatePaperRound(hand string) int {
	switch hand {
	case "A":
		return 6
	case "B":
		return 3
	case "C":
		return 0
	default:
		return 0
	}
}
