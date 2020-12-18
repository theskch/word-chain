package chain

import (
	"bytes"
	"container/list"
	"encoding/gob"
	"fmt"
)

// Chain represents a path of words. It stores all the words that were traveresed during the search and is responsible for word chain representation of the result.
type Chain struct {
	Links   []Link
	Visited map[string]bool
}

// Length returns the length of the chain (number of words in chain).
func (c Chain) Length() int {
	return len(c.Links)
}

// String returns the string representation of the chain
func (c Chain) String() string {
	var retVal string
	for i, link := range c.Links {
		retVal += link.Word
		if i != len(c.Links)-1 {
			retVal += "->"
		}
	}
	return retVal
}

// DeepCopy creates exact copy of `Chain` struct.
//
// Returns the copy of `Chain` struct or empty `Chain` struct and error
// if it is not possible to create a copy.
func (c Chain) DeepCopy() (Chain, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(c)
	if err != nil {
		return Chain{}, err
	}
	var copy Chain
	err = dec.Decode(&copy)
	return copy, err
}

// FindShortestChain returnes one of the shortest paths from `start` to `stop` word using dictionary `dict`.
// It uses the Breadth-first search algorithm for traversing a graph data structure. Nodes of the graph are
// words in the dictionary. Edges are created between two nodes with distance of 1 (same length words with one letter
// difference, case sensitive). Nodes of the graph are represented with `Link` structure, while edges are Conns
// inside `Link` structure.
//
// `start` is the first word in the chain.
//
// `stop` is the last word in the chain.
//
// `dict` is the dictionary from which the chain is created.
//
// Return value is one of the shortest chains. If there is no possible chain, empty chain is returned.
//
// Error is returned if len(start) != len(stop) or start/stop word is not present in the `dict`.
func FindShortestChain(start string, stop string, dict map[string]bool) (Chain, error) {
	// return error if length of the start and stop word doesn't match
	if len(start) != len(stop) {
		return Chain{}, fmt.Errorf("first and last words have different lengts")
	}

	links := make(map[string]Link)
	for key := range dict {
		links[key] = CreateLink(key, dict)
	}

	startLink, ok := links[start]
	// return error if start word is not present in the dictionary
	if !ok {
		return Chain{}, fmt.Errorf("first word is not present in the dictionary")
	}

	// return error if stop word is not present in the dictionary
	if _, ok = links[stop]; !ok {
		return Chain{}, fmt.Errorf("last word is not present in the dictionary")
	}

	queue := list.New()
	// first element in chain is the link with start word
	queue.PushBack(Chain{[]Link{startLink}, map[string]bool{start: true}})
	for queue.Len() > 0 {
		// take the first element of the queue and remove it from the queue.
		current := queue.Front()
		queue.Remove(current)
		chain := current.Value.(Chain)
		for _, conn := range chain.Links[len(chain.Links)-1].Conns {
			// if the word is aleary in chain, skip it. We don't want circular chains
			if chain.Visited[conn] {
				continue
			}
			// for each connection a new variation of chain is created
			chainCopy, err := chain.DeepCopy()
			if err != nil {
				continue
			}

			newLink := links[conn]
			chainCopy.Links = append(chainCopy.Links, newLink)
			chainCopy.Visited[newLink.Word] = true
			// if new link is made of the stop word, return the chain
			if newLink.Word == stop {
				return chainCopy, nil
			}

			// push back the new variation of chain, and continue the process
			queue.PushBack(chainCopy)
		}
	}

	// result is not found, return empty chain
	return Chain{}, nil
}
