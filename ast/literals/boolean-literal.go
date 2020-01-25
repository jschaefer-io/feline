package literals

type BooleanLiteral struct {
	literalBase
	Value bool
}

func (b *BooleanLiteral) ToString() string {
	if b.Value {
		return "true"
	}
	return "false"
}

func (b *BooleanLiteral) Get() (interface{}, error) {
	return b.Value, nil
}

func NewBooleanLiteral(val bool) BooleanLiteral {
	return BooleanLiteral{
		literalBase: literalBase{String},
		Value:       val,
	}
}
