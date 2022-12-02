package main

import "testing"

func TestParseData(t *testing.T) {
	data := "A Y\nA Y\nB X\nA Y\n"
	rounds := parseData(data)
	if len(rounds) != 4 {
		t.Fatalf("Wrong length returned expected 4 got %v\n", len(rounds))
	}

	round := rounds[0]
	if round[0] != "A" || round[1] != "Y" {
		t.Fatal("Wrong elements returned")
	}
}
