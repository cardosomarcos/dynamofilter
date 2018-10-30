package dynamofilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	eq := parser(ExpressionEquals)(name, value1)
	cn := parser(ExpressionContains)(city, cities)
	bw := parser(ExpressionBetween)(name, cities)
	in := parser(ExpressionIn)(name, cities)
	nt := parser(ExpressionNotIn)(name, cities)

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, in, "'name' IN (?,?)")
		assert.Equal(t, nt, "NOT ('name' IN (?,?))")
		assert.Equal(t, eq, "'name' = ?")
		assert.Equal(t, cn, "contains('city', ?,?)")
		assert.Equal(t, bw, "'name' BETWEEN ? AND ?")
	})
}
