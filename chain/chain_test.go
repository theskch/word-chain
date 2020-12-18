package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	t.Run("chain length", func(t *testing.T) {
		chain := Chain{
			Links: []Link{{}, {}, {}},
		}
		assert.Equal(t, 3, chain.Length())
	})
	t.Run("chain print", func(t *testing.T) {
		chain := Chain{
			Links: []Link{{Word: "one"}, {Word: "two"}, {Word: "three"}},
		}
		assert.Equal(t, "one->two->three", chain.String())
	})
	t.Run("chain deep copy", func(t *testing.T) {
		chain := Chain{
			Links:   []Link{{Word: "one", Conns: []string{"two"}}},
			Visited: map[string]bool{"one": true},
		}
		chainCopy, err := chain.DeepCopy()
		assert.NoError(t, err)
		assert.Equal(t, chain.Length(), chainCopy.Length())
		assert.Equal(t, chain.String(), chainCopy.String())
	})
}

func TestFindShortestChain(t *testing.T) {
	t.Run("start and stop word lenght mismatch", func(t *testing.T) {
		_, err := FindShortestChain("ruby", "cod", map[string]bool{})
		assert.Error(t, err)
	})
	t.Run("start word missing from the dictionary", func(t *testing.T) {
		_, err := FindShortestChain("ruby", "code", map[string]bool{"code": true})
		assert.Error(t, err)
	})
	t.Run("stop word missing from the dictionary", func(t *testing.T) {
		_, err := FindShortestChain("ruby", "code", map[string]bool{"ruby": true})
		assert.Error(t, err)
	})
	t.Run("no chains in dictionary", func(t *testing.T) {
		chain, err := FindShortestChain("ruby", "code", map[string]bool{"ruby": true, "code": true, "rube": true, "rode": true})
		assert.NoError(t, err)
		assert.Equal(t, "", chain.String())
		assert.Equal(t, 0, chain.Length())
	})
	t.Run("shortest chain found", func(t *testing.T) {
		chain, err := FindShortestChain("ruby", "code", map[string]bool{"ruby": true, "code": true, "rube": true, "rode": true, "mode": true, "rude": true})
		assert.NoError(t, err)
		assert.Equal(t, "ruby->rube->rude->rode->code", chain.String())
		assert.Equal(t, 5, chain.Length())
	})
}
