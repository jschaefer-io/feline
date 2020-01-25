package literals

import (
	"errors"
	"fmt"
)

type LiteralType uint8

const (
	Number           LiteralType = 0
	String           LiteralType = 1
	Boolean          LiteralType = 3
	Char             LiteralType = 4
	BinaryExpression LiteralType = 100
	Unknown          LiteralType = 255
)

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

func NewLiteral(literalType LiteralType, value interface{}) (Literal, error) {
	switch literalType {
	case Number:
		numLiteral := NewNumberLiteral(value.(float64))
		return &numLiteral, nil
	case Boolean:
		boolLiteral := NewBooleanLiteral(value.(bool))
		return &boolLiteral, nil
	case Char:
		charLiteral := NewCharLiteral(value.(rune))
		return &charLiteral, nil
	case String:
		strLiteral := NewStringLiteral(value.(string))
		return &strLiteral, nil
	}
	return nil, errors.New(fmt.Sprintf("unable to create literal from type %d", literalType))
}
