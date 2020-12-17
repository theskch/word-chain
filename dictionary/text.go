package dictionary

import (
	"bufio"
	"os"
)

// NewTextDictionary returnes new instance of Text dictionary populated with data from the file located in `path`.
func NewTextDictionary(path string) (Text, error) {
	dict := Text{
		words: map[string]bool{},
	}

	return dict, dict.load(path)
}

// Text dictionary uses a txt file to get the dictionary data into the map.
type Text struct {
	words map[string]bool
}

// Contains checks if dictionary contains `word`
func (t *Text) Contains(word string) bool {
	return t.words[word]
}

// GetWords returns words from the dictionary with specified length
//
// If `length` <= 0 all words are returned
func (t *Text) GetWords(length int) map[string]bool {
	if length <= 0 {
		return t.words
	}

	retVal := make(map[string]bool)
	for key := range t.words {
		if len(key) == length {
			retVal[key] = true
		}
	}

	return retVal
}

func (t *Text) load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t.words[scanner.Text()] = true
	}

	return scanner.Err()
}
