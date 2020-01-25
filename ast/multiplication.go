package ast

import (
	"errors"
	"fmt"
)

type Multiplication struct{}

func (m Multiplication) Calculate(exp *BinaryExpression) (interface{}, error) {
	typeA := exp.A.GetType()
	typeB := exp.B.GetType()
	if typeA != Number || typeB != Number {
		return 0, errors.New(fmt.Sprintf("operator not defined for type %d on %d", typeA, typeB))
	}
	valA, _ := exp.A.Get()
	valB, _ := exp.B.Get()
	return valA.(float64) * valB.(float64), nil
}
