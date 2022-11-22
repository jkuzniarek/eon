package ast

import (
	tk "eon/token"
	"bytes"
)

// ast node
type Node interface {
	TokenLiteral() string
	String() string
}
