package chain

import (
	"unicode/utf8"
)

// Link represents one word in chain
//
// It contains a string representation of the `word` and a map
// of the words with distance of 1 from the `word` (one letter appart)
type Link struct {
	Word  string
	Conns []string
}

// CreateLink returns a new link.
//
// Link is created from the `word` using dictionary `dict`.
//
// Comparison of words is case sensitive.
func CreateLink(word string, dict map[string]bool) Link {
	wordRunes := []rune(word)
	distance := 0
	conns := make([]string, 0)
	for key := range dict {
		// if key doesn't match the word in lenght, skip it
		if utf8.RuneCountInString(word) != utf8.RuneCountInString(key) {
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
			conns = append(conns, key)
		}
	}
	return Link{Word: word, Conns: conns}
}
