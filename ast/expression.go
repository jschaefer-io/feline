package ast

import "github.com/jschaefer-io/feline/ast/literals"

type Expression interface {
	literals.Literal
}