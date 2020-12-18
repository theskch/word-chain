package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
	t.Run("create link", func(t *testing.T) {
		dict := map[string]bool{"gold": true, "bold": true, "old": true, "load": true, "told": true}
		link := CreateLink("gold", dict)
		assert.Equal(t, "gold", link.Word)
		assert.Equal(t, 2, len(link.Conns))
		containsBold := false
		containsTold := false
		for _, val := range link.Conns {
			if val == "bold" {
				containsBold = true
			} else if val == "told" {
				containsTold = true
			}
		}
		assert.True(t, containsBold)
		assert.True(t, containsTold)
	})
}
