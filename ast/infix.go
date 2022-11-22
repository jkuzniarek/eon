package ast

import (
	tk "eon/token"
	"bytes"
)

type Infix struct {
	Token tk.Token // the operator token, eg +
	Left Node
	Operator string
	Right Node
}

func (ie *Infix) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *Infix) String() string {
	var out bytes.Buffer

	// out.WriteString("(")
	out.WriteString(ie.Left.String())
	if ie.Token.Type == tk.ACCESS_OPERATOR {
		out.WriteString(ie.Operator)
	}else	if ie.Token.Type == tk.ASSIGN_OPERATOR {
		out.WriteString(ie.Operator + " ")
	}else {
		out.WriteString(" " + ie.Operator + " ")
	}
	out.WriteString(ie.Right.String())
	// out.WriteString(")")

	return out.String()
}
