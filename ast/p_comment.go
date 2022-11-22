package ast

import (
	tk "eon/token"
	"bytes"
)

// comment
type Comment struct{
	Token tk.Token 
	Value string
	Multiline bool
}

func (n *Comment) TokenLiteral() string {
	return n.Token.Literal
}
func (n *Comment) String() string{
	var out bytes.Buffer
	if n.Multiline {
		out.WriteString("/*")
		out.WriteString(n.Value)
		out.WriteString("*/")
	}else{
		out.WriteString("//")
		out.WriteString(n.Value)
		out.WriteString("\n")
	}
	return out.String()
}