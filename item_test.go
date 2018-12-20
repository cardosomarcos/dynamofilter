package dynamofilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	equals := parser(ExpressionEquals)(name, value1)
	contains := parser(ExpressionContains)(city, cities)
	between := parser(ExpressionBetween)(name, cities)
	in := parser(ExpressionIn)(name, cities)
	notIn := parser(ExpressionNotIn)(name, cities)
	beginsWith := parser(ExpressionBeginsWith)(name, value1)
	less := parser(ExpressionLess)(name, value1)
	greater := parser(ExpressionGreater)(name, value1)
	lessOrEquals := parser(ExpressionLessOrEqual)(name, value1)
	greaterOrEquals := parser(ExpressionGreaterOrEqual)(name, value1)

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, in, "'name' IN (?,?)")
		assert.Equal(t, notIn, "NOT ('name' IN (?,?))")
		assert.Equal(t, equals, "'name' = ?")
		assert.Equal(t, contains, "contains('city', ?,?)")
		assert.Equal(t, between, "'name' BETWEEN ? AND ?")
		assert.Equal(t, beginsWith, "begins_with('name',?)")
		assert.Equal(t, less, "'name' < ?")
		assert.Equal(t, greater, "'name' > ?")
		assert.Equal(t, lessOrEquals, "'name' <= ?")
		assert.Equal(t, greaterOrEquals, "'name' >= ?")
	})
}
