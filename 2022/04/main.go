package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/danilomarques1/advent_of_code/2022/04/firstpuzzle"
	"github.com/danilomarques1/advent_of_code/2022/04/secondpuzzle"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := strings.Split(string(b), "\n")
	rowsSlice, err := parseData(rows)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("First puzzle = %v\n", firstpuzzle.Solve(rowsSlice))
	log.Printf("Second puzzle = %v\n", secondpuzzle.Solve(rowsSlice))
}

func parseData(rows []string) ([][][]int, error) {
	rowsSlice := make([][][]int, 0)
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		rowSlice, err := parseRow(row)
		if err != nil {
			return nil, err
		}
		rowsSlice = append(rowsSlice, rowSlice)
	}
	return rowsSlice, nil
}

func parseRow(row string) ([][]int, error) {
	rowSlice := make([][]int, 2)

	pairs := strings.Split(row, ",")
	if len(pairs) < 2 {
		return nil, errors.New("Error when spliting row")
	}

	firstPair := pairs[0]
	secondPair := pairs[1]

	firstPairSlice, err := convertPairToSlice(firstPair)
	if err != nil {
		return nil, err
	}
	rowSlice[0] = firstPairSlice

	secondPairSlice, err := convertPairToSlice(secondPair)
	if err != nil {
		return nil, err
	}
	rowSlice[1] = secondPairSlice

	return rowSlice, nil
}

func convertPairToSlice(pair string) ([]int, error) {
	pairSlice := make([]int, 2)
	pairSplit := strings.Split(pair, "-")
	if len(pairSplit) < 2 {
		return nil, errors.New("Error when spliting pair")
	}
	firstPairValue, err := strconv.Atoi(pairSplit[0])
	if err != nil {
		return nil, err
	}
	secondPairValue, err := strconv.Atoi(pairSplit[1])
	if err != nil {
		return nil, err
	}
	pairSlice[0] = firstPairValue
	pairSlice[1] = secondPairValue
	return pairSlice, nil
}
