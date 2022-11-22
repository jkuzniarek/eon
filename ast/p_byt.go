package ast

import (
	tk "eon/token"
	"bytes"
)

// byt
type Byt struct{
	Token tk.Token 
	Value []byte
}

func (il *Byt) TokenLiteral() string {
	return il.Token.Literal
}
func (il *Byt) String() string{
	return il.Token.Literal
}
