package ast

import (
	tk "eon/token"
	"bytes"
)

type Name struct {
	Token tk.Token // usually the tk.NAME token
	Value string
}
func (i *Name) TokenLiteral() string {
	return i.Token.Literal 
}
func (i *Name) String() string{
	return i.Value
}
