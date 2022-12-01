package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstPuzzle()
}

func secondPuzzle() {
}

func firstPuzzle() {
	data, err := readFileData("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	elves, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}
	position := getElfPositionWithHighestCalorie(elves)

	log.Println(position)
	log.Println(elves[position])

}

func readFileData(fileName string) ([]string, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), "\n"), nil
}

// will be returning the position of the elf and the amount of calorie hes carrying
func parseData(data []string) (map[int]int, error) {
	elves := make(map[int]int)
	position := 1
	totCal := 0
	for i, row := range data {
		if row == "" || i == len(data) - 1 {
			elves[position] = totCal
			position += 1
			totCal = 0
		} else {
			cal, err := strconv.Atoi(row)
			if err != nil {
				return nil, err
			}
			totCal += cal
		}
	}
	return elves, nil
}

func getElfPositionWithHighestCalorie(elves map[int]int) int {
	biggest := 0
	biggestPosition := 0
	for position, cal := range elves {
		if cal > biggest {
			biggestPosition = position
			biggest = cal
		}
	}
	return biggestPosition
}
