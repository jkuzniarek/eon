package ast

import (
	tk "eon/token"
)

// unsigned int
type UInt struct{
	Token tk.Token 
	Value uint
}

func (il *UInt) TokenLiteral() string {
	return il.Token.Literal
}
func (il *UInt) String() string{
	return il.Token.Literal
}
