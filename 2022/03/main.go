package main

import (
	"log"
	"os"
	"strings"

	"github.com/danilomarques1/advent_of_code/2022/03/firstpuzzle"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	rucksacks := strings.Split(string(b), "\n")
	log.Println(firstpuzzle.Solve(rucksacks))
}
