package filter

import (
	"pretty"
)

type Item struct {
	Property   string      `json:"property"`
	Value      interface{} `json:"value"`
	Query      string      `json:"query"`
	Expression Expression  `json:"expression"`
}

func itemBuilder(property string, value interface{}, expression Expression) Item {
	return Item{
		Property:   property,
		Value:      value,
		Query:      parser(expression)(property, value),
		Expression: expression,
	}
}

func parser(exp Expression) func(property string, value interface{}) string {
	var v string
	switch exp {
	case ExpressionContains:
		v = "contains('%s', %s)"
	case ExpressionEquals:
		v = "'%s' = %s"
	case ExpressionIn:
		v = "'%s' IN (%s)"
	}

	return func(property string, value interface{}) string {
		return pretty.Sprintf(v, property, repeat(value))
	}
}
