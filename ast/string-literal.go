package ast

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

func NewStringLiteral(content string) StringLiteral {
	return StringLiteral{
		literalBase: literalBase{String},
		Value:       content,
	}
}