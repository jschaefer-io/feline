package literals

import "strconv"

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

func NewNumberLiteral(number float64) NumberLiteral {
	return NumberLiteral{
		literalBase: literalBase{Number},
		Value:       number,
	}
}