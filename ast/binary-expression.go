package ast

import (
	"github.com/jschaefer-io/feline/ast/literals"
	"github.com/jschaefer-io/feline/ast/operators"
)

type BinaryExpression struct {
	Operator operators.Operator
	A        literals.Literal
	B        literals.Literal
}

func (b *BinaryExpression) Get() (interface{}, error) {
	return b.Operator.Calculate(b.A, b.B)
}
