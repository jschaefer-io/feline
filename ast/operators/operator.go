package operators

import (
	"github.com/jschaefer-io/feline/ast/literals"
)

// Operator
type Operator interface {
	Calculate(a literals.Literal, b literals.Literal) (interface{}, error, literals.LiteralType)
}
