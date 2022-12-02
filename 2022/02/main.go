package main

import (
	"log"
	"os"
	"strings"

	//"github.com/danilomarques1/advent_of_code/2022/02/firstpuzzle"
	"github.com/danilomarques1/advent_of_code/2022/02/secondpuzzle"
)

func main() {
	str, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rounds := parseData(string(str))
	//firstPuzzleTotalScore := firstpuzzle.Solve(rounds)
	//log.Printf("First puzzle = %v\n", firstPuzzleTotalScore)
	secondPuzzleTotalScore := secondpuzzle.Solve(rounds)
	log.Printf("Second puzzle = %v\n", secondPuzzleTotalScore)
}

func parseData(str string) [][]string {
	rounds := make([][]string, 0)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		round := strings.Split(line, " ")
		rounds = append(rounds, round)
	}
	return rounds
}
