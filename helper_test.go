package dynamofilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	c := repeat(cities)
	n := repeat(names)

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, c, "?,?")
		assert.Equal(t, n, "?,?,?")
	})
}
