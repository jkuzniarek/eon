package ast

import (
	tk "eon/token"
	"bytes"
)

// signed int
type SInt struct{
	Token tk.Token 
	Value int
}

func (il *SInt) TokenLiteral() string {
	return il.Token.Literal
}
func (il *SInt) String() string{
	return il.Token.Literal
}
