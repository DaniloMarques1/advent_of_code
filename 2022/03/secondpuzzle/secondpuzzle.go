package secondpuzzle

import "github.com/danilomarques1/advent_of_code/2022/03/firstpuzzle"

func Solve(grouped [][]string) int {
	totalPriority := 0
	for _, group := range grouped {
		totalPriority += groupTotalPriority(group)
	}

	return totalPriority
}

func groupTotalPriority(group []string) int {
	if len(group) < 3 {
		return 0
	}
	firstElfMap := buildMap(group[0])
	secondElfMap := buildMap(group[1])
	thirdElfMap := buildMap(group[2])
	return sumRepeated(firstElfMap, secondElfMap, thirdElfMap)
}

func buildMap(row string) map[string]int {
	m := make(map[string]int)
	for _, c := range row {
		m[string(c)] = 1
	}
	return m
}

func sumRepeated(firstElfMap, secondElfMap, thirdElfMap map[string]int) int {
	for key := range firstElfMap {
		if _, secondExist := secondElfMap[key]; secondExist {
			if _, thirdExist := thirdElfMap[key]; thirdExist {
				return firstpuzzle.ITEM_TYPE[key]
			}
		}
	}
	return 0
}
