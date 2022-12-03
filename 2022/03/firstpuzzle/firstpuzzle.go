package firstpuzzle

var ITEM_TYPE = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func Solve(rucksacks []string) int {
	totPriority := 0
	for _, rucksack := range rucksacks {
		firstcompartment, secondcompartment := getCompartmentFromRuckSack(rucksack)
		firstPartMapCompartment := buildMapFromCompartment(firstcompartment)
		secondPartMapCompartment := buildMapFromCompartment(secondcompartment)
		repeated := countOccurrences(firstPartMapCompartment, secondPartMapCompartment)
		totPriority += sumRepeatedItemsPriority(repeated)
	}

	return totPriority
}

func getCompartmentFromRuckSack(rucksack string) (string, string) {
	half := len(rucksack) / 2
	return rucksack[0:half], rucksack[half:]
}

func buildMapFromCompartment(compartment string) map[string]int {
	m := make(map[string]int)
	for _, c := range compartment {
		m[string(c)] = 1
	}
	return m
}

func countOccurrences(firstCompartmentMap map[string]int, secondCompartmentMap map[string]int) []string {
	repeated := make([]string, 0)
	for key := range firstCompartmentMap {
		if _, ok := secondCompartmentMap[key]; ok {
			repeated = append(repeated, key)
		}
	}
	return repeated
}

func sumRepeatedItemsPriority(repeated []string) int {
	totPriority := 0
	for _, item := range repeated {
		totPriority += ITEM_TYPE[string(item)]
	}
	return totPriority
}
