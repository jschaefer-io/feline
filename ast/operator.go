package ast

// Operator
type Operator interface {
	Calculate(exp *BinaryExpression) (interface{}, error)
}
