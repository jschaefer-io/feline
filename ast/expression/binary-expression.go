package expression

import (
	"github.com/jschaefer-io/feline/ast/literals"
	"github.com/jschaefer-io/feline/ast/operators"
)

type BinaryExpression struct {
	Operator operators.Operator
	A        literals.Literal
	B        literals.Literal
}

func (b *BinaryExpression) GetType() literals.LiteralType {
	return literals.BinaryExpression
}

func (b *BinaryExpression) ToString() string {
	panic("implement me")
}

func (b *BinaryExpression) Get() (literals.Literal, error) {
	res, err, resType := b.Operator.Calculate(b.A, b.B)
	if err != nil {
		return nil, err
	}
	return literals.NewLiteral(resType, res)
}
