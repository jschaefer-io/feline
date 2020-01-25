package literals

type CharLiteral struct {
	literalBase
	Value rune
}

func (char *CharLiteral) ToString() string {
	return string(char.Value)
}

func (char *CharLiteral) Get() (interface{}, error) {
	return char.Value, nil
}

func NewCharLiteral(char rune) CharLiteral {
	return CharLiteral{
		literalBase: literalBase{Char},
		Value:       char,
	}
}
