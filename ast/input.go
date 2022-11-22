package ast

import (
	"bytes"
)


type Input struct {
	Left Node
	Input Node
}

func (i *Input) TokenLiteral() string {
	return i.Left.TokenLiteral()
}
func (i *Input) String() string {
	var out bytes.Buffer

	out.WriteString(i.Left.String())
	out.WriteString(" ")
	out.WriteString(i.Input.String())
	
	return out.String()
}
