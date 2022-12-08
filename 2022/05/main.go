package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	//"github.com/danilomarques1/advent_of_code/2022/05/firstpuzzle"
	"github.com/danilomarques1/advent_of_code/2022/05/secondpuzzle"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	stacks := map[int][]string{
		1: {"S", "C", "V", "N"},
		2: {"Z", "M", "J", "H", "N", "S"},
		3: {"M", "C", "T", "G", "J", "N", "D"},
		4: {"T", "D", "F", "J", "W", "R", "M"},
		5: {"P", "F", "H"},
		6: {"C", "T", "Z", "H", "J"},
		7: {"D", "P", "R", "Q", "F", "S", "L", "Z"},
		8: {"C", "S", "L", "H", "D", "F", "P", "W"},
		9: {"D", "S", "M", "P", "F", "N", "G", "Z"},
	}

	plays := parseData(lines)
	//first := firstpuzzle.Solve(stacks, plays)
	//log.Printf("First puzzle %v\n", first)
	second := secondpuzzle.Solve(stacks, plays)
	log.Printf("second puzzle %v\n", second)
}

func parseData(lines []string) [][]int {
	plays := make([][]int, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		play := make([]int, 0, 3)
		words := strings.Split(line, " ")
		for _, word := range words {
			if n, isNumeric := isNumericString(word); isNumeric {
				play = append(play, n)
			}
		}
		plays = append(plays, play)
	}

	return plays
}

func isNumericString(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return n, true
}
