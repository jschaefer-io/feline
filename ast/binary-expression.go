package ast

type BinaryExpression struct {
	Operator Operator
	A        Literal
	B        Literal
}

func (b *BinaryExpression) Get() (interface{}, error) {
	return b.Operator.Calculate(b)
}
