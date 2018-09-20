package dynamofilter

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	filter := NewFilter().Equals(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' = ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionEquals)
	})
}

func TestContains(t *testing.T) {
	filter := NewFilter().Contains(name, []string{value1, value2})

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "contains('name', ?,?)")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, []string{value1, value2})
		assert.Equal(t, res.Expression, ExpressionContains)
	})
}

func TestIn(t *testing.T) {
	filter := NewFilter().In(name, []string{value1, value2})

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' IN (?,?)")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, []string{value1, value2})
		assert.Equal(t, res.Expression, ExpressionIn)
	})
}

func TestBuilder(t *testing.T) {
	filter := NewFilter().Equals(name, value1).Contains(city, []string{cityValue1, cityValue2})

	t.Run("success", func(t *testing.T) {
		filter, args := filter.Builder()

		assert.Equal(t, filter, pretty.Sprintf("'%s' = ? AND contains('%s', ?,?)", name, city))
		assert.Equal(t, args, []interface{}{
			value1,
			[]string{cityValue1, cityValue2},
		})
	})
}
