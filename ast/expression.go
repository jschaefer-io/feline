package ast

import (
	"errors"
	"fmt"
	"strconv"
)

type Expression interface {
	Literal
}

type BinaryExpression struct {
	Operator Operator
	A        Literal
	B        Literal
}

func (b *BinaryExpression) Get() (interface{}, error) {
	return b.Operator.calculate(b)
}

// Operator
type Operator interface {
	calculate(exp *BinaryExpression) (interface{}, error)
}

type Addition struct{}

func (a Addition) calculate(exp *BinaryExpression) (interface{}, error) {
	if exp.A.GetType() == Number && exp.B.GetType() == Number {
		valA, _ := exp.A.Get()
		valB, _ := exp.B.Get()
		return valA.(float64) + valB.(float64), nil
	}
	return exp.A.ToString() + exp.B.ToString(), nil
}

type Multiplication struct{}

func (m Multiplication) calculate(exp *BinaryExpression) (interface{}, error) {
	typeA := exp.A.GetType()
	typeB := exp.B.GetType()
	if typeA != Number || typeB != Number {
		return 0, errors.New(fmt.Sprintf("operator not defined for type %d on %d", typeA, typeB))
	}
	valA, _ := exp.A.Get()
	valB, _ := exp.B.Get()
	return valA.(float64) * valB.(float64), nil
}

// Literal
type Literal interface {
	Get() (interface{}, error)
	GetType() LiteralType
	ToString() string
}

type literalBase struct {
	literalType LiteralType
}

func (l *literalBase) GetType() LiteralType {
	return l.literalType
}

type NumberLiteral struct {
	literalBase
	Value float64
}

func (n *NumberLiteral) ToString() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 64)
}

func (n *NumberLiteral) Get() (interface{}, error) {
	return n.Value, nil
}

type StringLiteral struct {
	literalBase
	Value string
}

func (s *StringLiteral) ToString() string {
	return s.Value
}

func (s StringLiteral) Get() (interface{}, error) {
	return s.Value, nil
}

// Literal Constructors

func NewNumberLiteral(number float64) NumberLiteral {
	return NumberLiteral{
		literalBase: literalBase{Number},
		Value:       number,
	}
}

func NewStringLiteral(content string) StringLiteral {
	return StringLiteral{
		literalBase: literalBase{String},
		Value:       content,
	}
}

// Literal Types
type LiteralType uint8

const (
	Number LiteralType = 0
	String LiteralType = 1
)
