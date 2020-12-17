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
		assert.True(t, link.Conns["bold"])
		assert.True(t, link.Conns["told"])
	})
}
