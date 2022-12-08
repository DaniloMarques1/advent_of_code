package secondpuzzle

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

	var elements []string
	from, elements = pop(from, play[0])
	where = append(where, elements...)

	stacks[play[1]] = from
	stacks[play[2]] = where

	return stacks
}

// remove n elements from the crate and returns the removed elements on a new slice
func pop(crate []string, n int) ([]string, []string) {
	if len(crate) == 0 {
		return []string{}, []string{}
	}
	lastIndex := len(crate)
	begingRemovingAt := lastIndex - n
	if begingRemovingAt < 0 {
		begingRemovingAt = lastIndex
	}
	removed := crate[begingRemovingAt:]
	crate = crate[:begingRemovingAt]

	return crate, removed
}
