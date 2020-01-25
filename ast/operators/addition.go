package operators

import (
	"github.com/jschaefer-io/feline/ast/literals"
)

type Addition struct{}

func (add *Addition) Calculate(a literals.Literal, b literals.Literal) (interface{}, error, literals.LiteralType) {
	if a.GetType() == literals.Number && b.GetType() == literals.Number {
		valA, _ := a.Get()
		valB, _ := b.Get()
		return valA.(float64) + valB.(float64), nil, literals.Number
	}
	return a.ToString() + b.ToString(), nil, literals.String
}
