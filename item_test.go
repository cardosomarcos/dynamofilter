package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	eq := parser(ExpressionEquals)(name, value1)
	cn := parser(ExpressionContains)(city, cities)

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, eq, "'name' = ?")
		assert.Equal(t, cn, "contains('city', ?,?)")
	})
}
