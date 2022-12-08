package main

import (
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Marker %v\n", findMarker(string(b)))
	log.Printf("message %v\n", findStartOfMessage(string(b)))
}

func findMarker(datastream string) int {
	return lookUpStoppingAt(datastream, 4)
}

func findStartOfMessage(datastream string) int {
	return lookUpStoppingAt(datastream, 14)
}

func lookUpStoppingAt(datastream string, stopAt int) int {
	for i := 0; i < len(datastream); i++ {
		buffer := make([]rune, 0)
		buffer = append(buffer, rune(datastream[i]))
		for j := i + 1; j < len(datastream); j++ {
			if contains(rune(datastream[j]), buffer) {
				break
			}
			buffer = append(buffer, rune(datastream[j]))
			if len(buffer) == stopAt {
				return j + 1
			}
		}
	}
	return 0
}

func contains(c rune, s []rune) bool {
	for _, element := range s {
		if element == c {
			return true
		}
	}
	return false
}
