package main

import (
	"log"
	"os"
	"strings"

	"github.com/danilomarques1/advent_of_code/2022/03/firstpuzzle"
	"github.com/danilomarques1/advent_of_code/2022/03/secondpuzzle"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := strings.Split(string(b), "\n")
	log.Println(firstpuzzle.Solve(rucksacks))

	grouped := parseDataForSecondPuzzle(rucksacks)
	if grouped != nil {
		log.Println(secondpuzzle.Solve(grouped))
	}
}

func parseDataForSecondPuzzle(rucksacks []string) [][]string {
	if len(rucksacks) < 3 {
		return nil
	}

	grouped := make([][]string, 0)
	group := make([]string, 0)
	for _, rucksack := range rucksacks {
		group = append(group, rucksack)
		if len(group) == 3 {
			grouped = append(grouped, group)
			group = make([]string, 0)
		}
	}

	return grouped
}
