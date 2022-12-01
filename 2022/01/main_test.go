package main

import (
	"os"
	"testing"
)

func TestReadFileData(t *testing.T) {
	fileName := "temp.txt"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Coult not create file %v", err)
	}
	defer os.Remove(fileName)
	defer file.Close()

	str := "10\n10\n20\n\n20\n20\n"
	_, err = file.WriteString(str)
	if err != nil {
		t.Fatalf("Coult not write to file %v", err)
	}

	data, err :=  readFileData(fileName)
	if err != nil {
		t.Fatalf("Coult not write to file %v", err)
	}

	if len(data) != 7 {
		t.Fatalf("Should have 7 rows only got %v", len(data))
	}
}

func TestParseData(t *testing.T) {
	data := []string{"10", "20", "30", "", "10", "5", ""}
	elves, err := parseData(data)
	if err != nil {
		t.Fatalf("Coult not create file %v", err)
	}

	if len(elves) != 2 {
		t.Fatalf("Should have two keys got %v\n", len(elves))
	}

	if elves[60] != 1 {
		t.Fatalf("Should have 60 got %v\n", elves[1])
	}
}

func TestGetSortedCalories(t *testing.T) {
	elves := map[int]int{200: 2, 120: 3, 500: 1, 70: 4}
	sorted := getSortedCalories(elves)
	if sorted[0] != 500 {
		t.Fatalf("Wrong first element, should be 500 got %v\n", sorted[0])
	}
}

func TestGetTotalCaloriesFromTopThree(t *testing.T) {
	sorted := []int{200, 100, 50, 20}
	total := getTotalCaloriesFromTopThree(sorted)
	if total != 350 {
		t.Fatalf("Wrong total, should be 350 got %v\n", total)
	}
}
