package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)


// find how much calorie an elf is carrying
func main() {
	data, err := readFileData("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	elves, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}
	caloriesSorted := getSortedCalories(elves)

	firstElfTotalCalorie := getElfPositionWithHighestCalorie(caloriesSorted)

	log.Println(elves[firstElfTotalCalorie])
	log.Println(firstElfTotalCalorie)

	topThreeTotal := getTotalCaloriesFromTopThree(caloriesSorted)
	log.Println(topThreeTotal)
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
			elves[totCal] = position
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

func getSortedCalories(elves map[int]int) []int {
	calories := make([]int, len(elves))
	i := 0
	for cal := range elves {
		calories[i] = cal
		i += 1
	}

	sort.SliceStable(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return calories
}

func getElfPositionWithHighestCalorie(calories []int) int {
	return calories[0]
}

func getTotalCaloriesFromTopThree(calories []int) int {
	if len(calories) < 3 {
		return 0
	}

	first := calories[0]
	second := calories[1]
	third := calories[2]

	return first + second + third
}
