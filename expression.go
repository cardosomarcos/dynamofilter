package dynamofilter

type Expression string

const (
	ExpressionContains Expression = "contains('%s', %s)"
	ExpressionEquals   Expression = "'%s' = %s"
	ExpressionIn       Expression = "'%s' IN (%s)"
	ExpressionNotIn    Expression = "NOT ('%s' IN (%s))"
	ExpressionBetween  Expression = "'%s' BETWEEN ? AND ?%.s"
)

func (exp Expression) toString() string {
	return string(exp)
}
