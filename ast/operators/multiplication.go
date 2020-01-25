package operators

import (
	"errors"
	"fmt"
	"github.com/jschaefer-io/feline/ast/literals"
)

type Multiplication struct{}

func (mul *Multiplication) Calculate(a literals.Literal, b literals.Literal) (interface{}, error) {
	typeA := a.GetType()
	typeB := b.GetType()
	if typeA != literals.Number || typeB != literals.Number {
		return 0, errors.New(fmt.Sprintf("operator not defined for type %d on %d", typeA, typeB))
	}
	valA, _ := a.Get()
	valB, _ := b.Get()
	return valA.(float64) * valB.(float64), nil
}
