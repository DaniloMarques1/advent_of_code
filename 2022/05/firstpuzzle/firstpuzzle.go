package firstpuzzle

func Solve(stacks map[int][]string, plays [][]int) string {
	stacks = Play(stacks, plays)
	return getTopCrates(stacks)
}

func getTopCrates(stacks map[int][]string) string {
	answer := ""
	for i := 1; i <= len(stacks); i++ {
		crates := stacks[i]
		if len(crates) > 0 {
			last := crates[len(crates)-1]
			answer += last
		}
	}
	return answer
}

func Play(stacks map[int][]string, plays [][]int) map[int][]string {
	for _, play := range plays {
		stacks = rearrange(stacks, play)
	}
	return stacks
}

// play[0] tell us how many to move
// play[1] tell us from where to move
// play[2] tell us where to move
func rearrange(stacks map[int][]string, play []int) map[int][]string {
	from := stacks[play[1]]
	where := stacks[play[2]]

	for i := 0; i < play[0]; i++ {
		var element string
		from, element = pop(from)
		if len(element) > 0 {
			where = append(where, element)
		}
	}

	stacks[play[1]] = from
	stacks[play[2]] = where

	return stacks
}

// remove and return the last element
func pop(crate []string) ([]string, string) {
	if len(crate) == 0 {
		return []string{}, ""
	}

	lastIndex := len(crate) - 1
	lastElement := crate[lastIndex]
	crate = crate[0:lastIndex]

	return crate, lastElement
}
