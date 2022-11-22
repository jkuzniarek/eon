package ast

import (
	tk "eon/token"
	"bytes"
)

// str
type Str struct{
	Token tk.Token 
	Value string
}

func (il *Str) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Str) String() string{
	return il.Token.Literal
}
