package literals

type LiteralType uint8

const (
	Number  LiteralType = 0
	String  LiteralType = 1
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
