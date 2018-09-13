package filter

type Expression string

const (
	ExpressionContains Expression = "contains"
	ExpressionEquals   Expression = "equals"
	ExpressionIn       Expression = "in"
)
