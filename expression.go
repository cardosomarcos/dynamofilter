package dynamofilter

type Expression string

const (
	ExpressionContains Expression = "CONTAINS"
	ExpressionEquals   Expression = "EQUALS"
	ExpressionIn       Expression = "IN"
	ExpressionNotIn    Expression = "NOT_IN"
)
