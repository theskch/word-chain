package chain

// Link represents one word in chain
//
// It contains a string representation of the `word` and a map
// of the words with distance of 1 from the `word` (one letter appart)
type Link struct {
	Word  string
	Conns map[string]bool
}

// CreateLink returns a new link.
//
// Link is created from the `word` using dictionary `dict`.
//
// Comparison of words is case sensitive.
func CreateLink(word string, dict map[string]bool) Link {
	wordRunes := []rune(word)
	distance := 0
	conns := make(map[string]bool)

	for key := range dict {
		// if key doesn't match the word in lenght, skip it
		if len(word) != len(key) {
			continue
		}

		distance = 0
		for i, rk := range []rune(key) {
			if wordRunes[i] != rk {
				distance++
			}

			if distance > 1 {
				break
			}
		}

		// don't include the `word` in connections
		if distance == 1 {
			conns[key] = true
		}
	}
	return Link{Word: word, Conns: conns}
}
