package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTextDictionary(t *testing.T) {
	t.Run("wrong path", func(t *testing.T) {
		_, err := NewTextDictionary("dictionary_large.txt")
		assert.Error(t, err)
	})
	t.Run("load small dictionary", func(t *testing.T) {
		dict, err := NewTextDictionary("dictionary_small.txt")
		assert.NoError(t, err)
		assert.True(t, dict.Contains("ruby"))
		assert.True(t, dict.Contains("rubs"))
		assert.True(t, dict.Contains("robs"))
		assert.True(t, dict.Contains("rods"))
		assert.True(t, dict.Contains("rode"))
		assert.True(t, dict.Contains("code"))
		assert.True(t, dict.Contains("lead"))
		assert.True(t, dict.Contains("load"))
		assert.True(t, dict.Contains("goad"))
		assert.True(t, dict.Contains("gold"))
		assert.True(t, dict.Contains("cat"))
		assert.True(t, dict.Contains("cot"))
		assert.True(t, dict.Contains("cog"))
		assert.True(t, dict.Contains("dog"))
	})
	t.Run("get words from dictionary", func(t *testing.T) {
		dict, _ := NewTextDictionary("dictionary_small.txt")
		subDict := dict.GetWords(0)
		assert.Equal(t, 14, len(subDict))
		assert.True(t, subDict["ruby"])
		assert.True(t, subDict["rubs"])
		assert.True(t, subDict["robs"])
		assert.True(t, subDict["rods"])
		assert.True(t, subDict["rode"])
		assert.True(t, subDict["code"])
		assert.True(t, subDict["lead"])
		assert.True(t, subDict["load"])
		assert.True(t, subDict["goad"])
		assert.True(t, subDict["gold"])
		assert.True(t, subDict["cat"])
		assert.True(t, subDict["cot"])
		assert.True(t, subDict["cog"])
		assert.True(t, subDict["dog"])

		subDict = dict.GetWords(3)
		assert.Equal(t, 4, len(subDict))
		assert.True(t, subDict["cat"])
		assert.True(t, subDict["cot"])
		assert.True(t, subDict["cog"])
		assert.True(t, subDict["dog"])

		subDict = dict.GetWords(5)
		assert.Equal(t, 0, len(subDict))
	})
}
