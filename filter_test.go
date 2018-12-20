package dynamofilter

import (
	"github.com/kr/pretty"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_Equals(t *testing.T) {
	filter := NewFilter().Equals(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' = ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionEquals)
	})
}

func TestFilter_Contains(t *testing.T) {
	filter := NewFilter().Contains(name, []string{value1, value2})

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "contains('name', ?,?)")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, []string{value1, value2})
		assert.Equal(t, res.Expression, ExpressionContains)
	})
}

func TestFilter_In(t *testing.T) {
	filter := NewFilter().In(name, []string{value1, value2})

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' IN (?,?)")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, []string{value1, value2})
		assert.Equal(t, res.Expression, ExpressionIn)
	})
}

func TestFilter_NotIn(t *testing.T) {
	filter := NewFilter().NotIn(name, []string{value1, value2})

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "NOT ('name' IN (?,?))")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, []string{value1, value2})
		assert.Equal(t, res.Expression, ExpressionNotIn)
	})
}

func TestFilter_Between(t *testing.T) {
	filter := NewFilter().Between(name, []string{value1, value2})

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' BETWEEN ? AND ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, []string{value1, value2})
		assert.Equal(t, res.Expression, ExpressionBetween)
	})
}

func TestFilter_BeginsWith(t *testing.T) {
	filter := NewFilter().BeginsWith(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "begins_with('name',?)")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionBeginsWith)
	})
}

func TestFilter_Less(t *testing.T) {
	filter := NewFilter().Less(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' < ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionLess)
	})
}

func TestFilter_Greater(t *testing.T) {
	filter := NewFilter().Greater(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' > ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionGreater)
	})
}

func TestFilter_LessOrEquals(t *testing.T) {
	filter := NewFilter().LessOrEquals(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' <= ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionLessOrEqual)
	})
}

func TestFilter_GreaterOrEquals(t *testing.T) {
	filter := NewFilter().GreaterOrEquals(name, value1)

	res := filter.Get("name")

	t.Run("success", func(t *testing.T) {
		assert.Equal(t, res.Query, "'name' >= ?")
		assert.Equal(t, res.Property, "name")
		assert.Equal(t, res.Value, value1)
		assert.Equal(t, res.Expression, ExpressionGreaterOrEqual)
	})
}

func TestBuilder(t *testing.T) {
	filter := NewFilter().Equals(name, value1)
	filter.Contains(city, []string{cityValue1, cityValue2})

	t.Run("success", func(t *testing.T) {
		filter, args := filter.Builder()
		assert.Equal(t, filter, pretty.Sprintf("'%s' = ? AND contains('%s', ?,?)", name, city))
		assert.Equal(t, args, []interface{}{value1, cityValue1, cityValue2})
	})
}
